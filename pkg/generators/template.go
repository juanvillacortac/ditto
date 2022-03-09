package generators

import (
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/juanvillacortac/rosetta/pkg/ast"
)

func templateHelpers(models ast.ModelMap, options GenerateConfig) template.FuncMap {
	return template.FuncMap{
		"HaveDefaultValue": func(p ast.ModelProp) bool { return p.DefaultValue != nil },
		"PropDefaultValue": func(p ast.ModelProp) string {
			if p.DefaultValue == nil {
				return ""
			}
			if helper, ok := options.Helpers[*p.DefaultValue]; ok {
				return helper
			}
			return *p.DefaultValue
		},
		"Models":   func() ast.ModelMap { return models },
		"GetModel": func(modelName string) *ast.Model { return models[modelName] },
		"GetModelDeps": func(modelName string) []*ast.Model {
			deps := make([]*ast.Model, 0)
			depsStr := models.GetModelDeps(modelName, nil)
			for _, str := range depsStr {
				m := models[str]
				deps = append(deps, m)
			}
			return deps
		},
		"GetNodeOption":    ast.GetNodeOption,
		"ToCamelCase":      strcase.ToCamel,
		"ToLowerCamelCase": strcase.ToLowerCamel,
		"ToKebabCase":      strcase.ToKebab,
		"ToSnakeCase":      strcase.ToSnake,
	}
}
