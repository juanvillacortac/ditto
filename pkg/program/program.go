package program

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/juanvillacortac/rosetta/pkg/ast"
	"github.com/juanvillacortac/rosetta/pkg/generators"
	"github.com/juanvillacortac/rosetta/pkg/parser/proto"
	"github.com/juanvillacortac/rosetta/pkg/parser/yaml"

	y "gopkg.in/yaml.v2"
)

type ProgramConfig struct {
	Definitions    generators.Definitions      `json:"definitions" yaml:"definitions"`
	SchemaFile     string                      `json:"schema" yaml:"schema"`
	OutputBasePath string                      `json:"output" yaml:"output"`
	Generators     []generators.GenerateConfig `json:"generators"`

	root *ast.RootNode
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

func (p *ProgramConfig) Generate(verbose bool) ([]generators.OutputFile, error) {
	if p.root == nil {
		return nil, fmt.Errorf("schema not loaded")
	}
	schemaPath, _ := filepath.Abs(path.Dir(p.SchemaFile))
	files := make([]generators.OutputFile, 0)
	for i, g := range p.Generators {
		gApplied := g.ApplyDefinitions(p.Definitions)
		fmt.Fprintf(os.Stdout, "[%d/%d] %s\n", i+1, len(p.Generators), g.Name)
		fs, err := generators.Generate(schemaPath, p.root, gApplied, verbose)
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
