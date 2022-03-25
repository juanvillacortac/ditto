package generators

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/juanvillacortac/ditto/pkg/ast"
)

type (
	TypesMap map[string]string

	Definitions struct {
		Types   map[string]string `json:"types"`
		Helpers map[string]string `json:"helpers"`
	}

	DefinitionsMap map[string]struct {
		Types   map[string]string `json:"types"`
		Helpers map[string]string `json:"helpers"`
	}
)

type OutputFile struct {
	Filename string
	Body     string
}

type GenerateConfig struct {
	Name     string `json:"name" yaml:"name"`
	Template string `json:"template,omitempty" yaml:"template,omitempty"`
	Output   string `json:"output" yaml:"output"`
	Ignore   string `json:"ignore" yaml:"ignore"`

	From    string            `json:"from" yaml:"from"`
	Types   map[string]string `json:"types" yaml:"types"`
	Helpers map[string]string `json:"helpers" yaml:"helpers"`
}

func (g GenerateConfig) ApplyDefinitions(definitions DefinitionsMap) GenerateConfig {
	clone := g
	if g.From == "" {
		return g
	}
	if len(clone.Types) == 0 {
		if t, ok := definitions[clone.From]; ok {
			clone.Types = t.Types
		}
	}
	if len(clone.Helpers) == 0 {
		if t, ok := definitions[clone.From]; ok {
			clone.Helpers = t.Helpers
		}
	}
	return clone
}

func AdaptModel(models ast.ModelMap, definitions Definitions) ast.ModelMap {
	clone := make(ast.ModelMap)
	buff, _ := json.Marshal(models)
	if err := json.Unmarshal(buff, &clone); err != nil {
		panic(err)
	}
	for k, m := range clone {
		for i, p := range m.Props {
			if t, ok := definitions.Types[p.Type]; ok {
				clone[k].Props[i].Type = t
			}
			if p.DefaultValue != nil {
				if h, ok := definitions.Helpers[*p.DefaultValue]; ok {
					clone[k].Props[i].DefaultValue = &h
				}
			}
		}
	}
	return clone
}

func Generate(root *ast.RootNode, config GenerateConfig, verbose bool) ([]OutputFile, error) {
	models := AdaptModel(root.Models, Definitions{
		Types:   config.Types,
		Helpers: config.Helpers,
	})
	t, err := createTemplate(config.Name, config.Template, models, config)
	if err != nil {
		return nil, err
	}
	ot, err := createTemplate(config.Name, config.Output, models, config)
	if err != nil {
		return nil, err
	}
	files := make([]OutputFile, 0)
	cnt := 0
	for _, m := range models {
		context := &TemplateContext{
			Root:  root,
			Model: m,
		}
		if val, ok := m.ModelOptions[config.Ignore]; ok && val == "true" {
			continue
		}
		if verbose {
			fmt.Fprintf(os.Stdout, "-> [%d/%d] Generating \"%s\"\n", cnt+1, len(models), m.Name())
		}

		body, err := execTemplate(t, context)
		if err != nil {
			return nil, err
		}
		output, err := execTemplate(ot, context)
		if err != nil {
			return nil, err
		}
		files = append(files, OutputFile{
			Filename: output,
			Body:     body,
		})
		cnt++
	}
	return files, nil
}
