package generators

import (
	_ "embed"
	"os"
	"text/template"

	"github.com/juanvillacortac/rosetta/pkg/model"
)

//go:embed typescript-template.ts.template
var ts []byte

func Generate(models model.ModelMap) {
	t, err := template.New("todos").Parse(string(ts))
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, models)
	if err != nil {
		panic(err)
	}
}
