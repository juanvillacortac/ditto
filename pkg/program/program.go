package program

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/juanvillacortac/rosetta/pkg/ast"
	"github.com/juanvillacortac/rosetta/pkg/generators"
	"github.com/yoheimuta/go-protoparser/v4"
	proto "github.com/yoheimuta/go-protoparser/v4/parser"
)

type Program struct {
	File       string                      `json:"file"`
	Generators []generators.GenerateConfig `json:"generators"`

	source *proto.Proto
	models ast.ModelMap
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

	got, err := protoparser.Parse(
		reader,
		protoparser.WithDebug(config.debug),
		protoparser.WithPermissive(config.debug),
	)
	if config.debug {
		gotJSON, err := json.MarshalIndent(got, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to marshal, err %v\n", err)
		}
		os.WriteFile("./.rosetta_proto_ast.json", gotJSON, os.ModePerm)
	}

	if err != nil {
		return fmt.Errorf("[Proto parsing error]: %v", err)
	}
	p.source = got
	models, err := ast.GetModelsFromProto(got)
	if err != nil {
		return fmt.Errorf("[Models parsing error]: %v", err)
	}
	p.models = models

	return nil
}

func (p *Program) Generate() ([]generators.OutputFile, error) {
	files := make([]generators.OutputFile, 0)
	for i, g := range p.Generators {
		fmt.Fprintf(os.Stdout, "[%d/%d] %s\n", i+1, len(p.Generators), g.Name)
		fs, err := generators.Generate(p.models, g)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}
	return files, nil
}
