package program

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/juanvillacortac/rosetta/pkg/ast"
	"github.com/juanvillacortac/rosetta/pkg/generators"
	"github.com/juanvillacortac/rosetta/pkg/parser/proto"
	"github.com/juanvillacortac/rosetta/pkg/parser/yaml"

	y "gopkg.in/yaml.v2"
)

type ProgramConfig struct {
	SchemaFile     string                    `json:"schema,omitempty" yaml:"schema,omitempty"`
	OutputBasePath string                    `json:"output,omitempty" yaml:"output,omitempty"`
	Definitions    generators.DefinitionsMap `json:"definitions" yaml:"definitions"`
	Generators     []GenerateConfig          `json:"generators"`

	root *ast.RootNode
}

type GenerateConfig struct {
	Name     string `json:"name" yaml:"name"`
	Template string `json:"template" yaml:"template"`
	Output   string `json:"output" yaml:"output"`

	From    string            `json:"from" yaml:"from"`
	Types   map[string]string `json:"types" yaml:"types"`
	Helpers map[string]string `json:"helpers" yaml:"helpers"`
}

func NewProgramFromConfigFile(reader *os.File) (*ProgramConfig, error) {
	buffer := bytes.Buffer{}
	buffer.ReadFrom(reader)
	if _, err := buffer.ReadFrom(reader); err != nil {
		return nil, err
	}
	p := &ProgramConfig{}
	ext := path.Ext(reader.Name())
	switch ext {
	case ".yml", ".yaml":
		if err := y.Unmarshal(buffer.Bytes(), p); err != nil {
			return nil, err
		}
	case ".json":
		if err := json.Unmarshal(buffer.Bytes(), p); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf(`unsupported file extension, expect ".json", ".yml" or ".yaml", got: %v`, ext)
	}
	return p, nil
}

func (p *ProgramConfig) Parse() error {
	reader, err := os.Open(p.SchemaFile)
	if err != nil {
		err = fmt.Errorf("failed to open %s, err %v", p.SchemaFile, err)
		return err
	}
	defer reader.Close()

	if err != nil {
		return fmt.Errorf("[Proto parsing error]: %v", err)
	}
	ext := path.Ext(reader.Name())
	var root *ast.RootNode
	switch ext {
	case ".yml", ".yaml":
		root, err = yaml.GetRootNodeFromYaml(reader)
		if err != nil {
			return fmt.Errorf("[Models parsing error]: %v", err)
		}
	case ".proto":
		root, err = proto.GetRootNodeFromProto(reader)
	default:
		return fmt.Errorf("schema file extension not allowed")
	}
	if err != nil {
		return fmt.Errorf("[Models parsing error]: %v", err)
	}
	p.root = root

	return nil
}

func (p *ProgramConfig) LoadGenerateConfigsWithTemplates(relativePath string) ([]generators.GenerateConfig, error) {
	configs := make([]generators.GenerateConfig, 0)
	for i := range p.Generators {
		g := p.Generators[i]
		reader, err := os.Open(path.Join(relativePath, g.Template))
		if err != nil {
			err = fmt.Errorf("failed to open %s, err %v", g.Template, err)
			return nil, err
		}
		defer reader.Close()

		buffer := bytes.Buffer{}
		if _, err := buffer.ReadFrom(reader); err != nil {
			return nil, err
		}

		gg := generators.GenerateConfig{
			Name:     g.Name,
			Template: buffer.String(),
			Output:   g.Output,
			From:     g.From,
			Types:    g.Types,
			Helpers:  g.Helpers,
		}

		gApplied := gg.ApplyDefinitions(p.Definitions)
		configs = append(configs, gApplied)
	}
	return configs, nil
}

func (p *ProgramConfig) Generate(relativePath string, verbose bool) ([]generators.OutputFile, error) {
	if p.root == nil {
		return nil, fmt.Errorf("schema not loaded")
	}
	files := make([]generators.OutputFile, 0)
	configs, _ := p.LoadGenerateConfigsWithTemplates(relativePath)
	for i, g := range configs {
		fmt.Fprintf(os.Stdout, "[%d/%d] %s\n", i+1, len(configs), g.Name)
		fs, err := generators.Generate(p.root, g, verbose)
		if err != nil {
			return nil, err
		}
		for i := range fs {
			f := &fs[i]
			f.Filename = path.Join(p.OutputBasePath, f.Filename)
		}
		files = append(files, fs...)
	}
	return files, nil
}
