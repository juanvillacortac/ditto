package generators

import (
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/juanvillacortac/ditto/pkg/ast"
)

func templateHelpers(models ast.ModelMap, options GenerateConfig) template.FuncMap {
	return template.FuncMap{
		"Models":     func() ast.ModelMap { return models },
		"Model":      func(modelName string) *ast.Model { return models[modelName] },
		"ModelDeps":  models.ModelDependencies,
		"NodeOption": ast.GetNodeOption,
		"NodeOptionOr": func(n ast.Node, optionName, fallback string) string {
			if v := ast.GetNodeOption(n, optionName); v != nil {
				return *v
			}
			return fallback
		},
		"CamelCase":      strcase.ToCamel,
		"LowerCamelCase": strcase.ToLowerCamel,
		"KebabCase":      strcase.ToKebab,
		"SnakeCase":      strcase.ToSnake,
		"Plural": func(str string) string {
			p := pluralize.NewClient()
			return p.Plural(str)
		},
		"Singular": func(str string) string {
			p := pluralize.NewClient()
			return p.Singular(str)
		},
	}
}

type TemplateContext struct {
	Root  *ast.RootNode
	Model *ast.Model
}

func createTemplate(name string, content string, models ast.ModelMap, options GenerateConfig) (*template.Template, error) {
	return template.New(name).Funcs(templateHelpers(models, options)).Parse(content)
}

func execTemplate(t *template.Template, context *TemplateContext) (string, error) {
	writer := &strings.Builder{}
	if err := t.Execute(writer, context); err != nil {
		return "", err
	}
	return writer.String(), nil
}
