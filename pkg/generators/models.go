package generators

import "github.com/juanvillacortac/rosetta/pkg/ast"

type OutputFile struct {
	Filename string
	Body     string
}

func GetModelDeps(m ast.Model, deps []string, models ast.ModelMap) []string {
	clone := make([]string, len(deps))
	copy(clone, deps)
	for _, p := range m.Props {
		if mm, ok := models[string(p.Type)]; ok {
			clone = append(clone, mm.ModelName)
		}
	}
	return clone
}
