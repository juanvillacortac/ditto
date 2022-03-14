package generators

import (
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/juanvillacortac/rosetta/pkg/ast"
)

func templateHelpers(models ast.ModelMap, options GenerateConfig) template.FuncMap {
	return template.FuncMap{
		"Models":         func() ast.ModelMap { return models },
		"Model":          func(modelName string) *ast.Model { return models[modelName] },
		"ModelDeps":      models.ModelDependencies,
		"NodeOption":     ast.GetNodeOption,
		"CamelCase":      strcase.ToCamel,
		"LowerCamelCase": strcase.ToLowerCamel,
		"KebabCase":      strcase.ToKebab,
		"SnakeCase":      strcase.ToSnake,
	}
}
