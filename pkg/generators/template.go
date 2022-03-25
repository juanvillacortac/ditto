package generators

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/juanvillacortac/ditto/pkg/ast"

	"github.com/Masterminds/sprig"
)

func templateHelpers(t *template.Template, models ast.ModelMap, options GenerateConfig) template.FuncMap {
	funcs := template.FuncMap{
		"ExecTmpl": func(name string, obj interface{}) string {
			buf := &bytes.Buffer{}
			err := t.ExecuteTemplate(buf, name, obj)
			if err != nil {
				panic(err)
			}
			return buf.String()
		},
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
	for k, v := range sprig.FuncMap() {
		funcs[k] = v
	}
	return funcs
}

type TemplateContext struct {
	Root  *ast.RootNode
	Model *ast.Model
}

func createTemplate(name string, content string, models ast.ModelMap, options GenerateConfig) (*template.Template, error) {
	t := template.New(name)
	return t.Funcs(templateHelpers(t, models, options)).Parse(content)
}

func execTemplate(t *template.Template, context *TemplateContext) (string, error) {
	writer := &strings.Builder{}
	if err := t.Execute(writer, context); err != nil {
		return "", err
	}
	return writer.String(), nil
}
