package program

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/juanvillacortac/rosetta/pkg/ast"
	"github.com/juanvillacortac/rosetta/pkg/generators"
	"github.com/juanvillacortac/rosetta/pkg/parser/proto"
	"github.com/juanvillacortac/rosetta/pkg/parser/yaml"
)

type Program struct {
	File       string                      `json:"file"`
	Generators []generators.GenerateConfig `json:"generators"`

	root *ast.RootNode
}

func NewProgramFromJson(reader io.Reader) (*Program, error) {
	buffer := bytes.Buffer{}
	buffer.ReadFrom(reader)
	if _, err := buffer.ReadFrom(reader); err != nil {
		return nil, err
	}
	p := &Program{}
	if err := json.Unmarshal(buffer.Bytes(), p); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Program) Parse(options ...Option) error {
	config := &ParseConfig{
		permissive: true,
	}
	for _, opt := range options {
		opt(config)
	}

	reader, err := os.Open(p.File)
	if err != nil {
		err = fmt.Errorf("failed to open %s, err %v", p.File, err)
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
		root, err = proto.GetRootNodeFromProto(reader, &proto.ParseConfig{
			Debug:      config.debug,
			Permissive: config.permissive,
		})
	default:
		return fmt.Errorf("schema file extension not allowed")
	}
	if err != nil {
		return fmt.Errorf("[Models parsing error]: %v", err)
	}
	p.root = root

	return nil
}

func (p *Program) Generate() ([]generators.OutputFile, error) {
	if p.root == nil {
		return nil, fmt.Errorf("schema not loaded")
	}
	files := make([]generators.OutputFile, 0)
	for i, g := range p.Generators {
		fmt.Fprintf(os.Stdout, "[%d/%d] %s\n", i+1, len(p.Generators), g.Name)
		fs, err := generators.Generate(p.root, g)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}
	return files, nil
}
