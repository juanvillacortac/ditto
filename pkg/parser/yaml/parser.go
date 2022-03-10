package yaml

import (
	"bytes"
	"fmt"
	"io"

	"github.com/juanvillacortac/rosetta/pkg/ast"
	"github.com/juanvillacortac/rosetta/pkg/utils"
	"gopkg.in/yaml.v2"
)

func GetRootNodeFromYaml(reader io.Reader) (*ast.RootNode, error) {
	var tree yaml.MapSlice
	buffer := bytes.Buffer{}
	if _, err := buffer.ReadFrom(reader); err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(buffer.Bytes(), &tree); err != nil {
		return nil, err
	}
	root := &ast.RootNode{
		RootOptions: make(ast.Options),
		Models:      make(ast.ModelMap),
	}
	t := []yaml.MapItem{}
	for _, item := range tree {
		t = append(t, item)
	}
	for _, val := range t {
		key := val.Key.(string)
		switch val.Value.(type) {
		case string:
			o, isOption := utils.UnwrapString(key, "(", ")")
			if isOption {
				root.RootOptions[o] = val.Value.(string)
			}
		case yaml.MapSlice:
			m := &ast.Model{
				ModelName:    key,
				Props:        make([]*ast.ModelProp, 0),
				ModelOptions: make(ast.Options),
			}
			for _, prop := range val.Value.(yaml.MapSlice) {
				switch prop.Value.(type) {
				case int, float32, float64:
					val := prop.Value
					if o, isOption := utils.UnwrapString(key, "(", ")"); isOption {
						m.ModelOptions[o] = fmt.Sprintf("%v", val)
					}
				case string:
					val := prop.Value.(string)
					if o, isOption := utils.UnwrapString(prop.Key.(string), "(", ")"); isOption {
						m.ModelOptions[o] = val
					} else {
						p := &ast.ModelProp{
							PropName:     prop.Key.(string),
							IsRequired:   true,
							IsArray:      false,
							DefaultValue: nil,
							Type:         val,
							PropOptions:  make(ast.Options),
						}
						m.Props = append(m.Props, p)
					}
				case yaml.MapSlice:
					val := prop.Value.(yaml.MapSlice)
					p := &ast.ModelProp{
						PropName:    prop.Key.(string),
						IsRequired:  true,
						PropOptions: make(ast.Options),
					}
					for _, meta := range val {
						if o, isOption := utils.UnwrapString(meta.Key.(string), "(", ")"); isOption {
							p.PropOptions[o] = fmt.Sprint(meta)
						} else {
							switch meta.Key.(string) {
							case "default":
								var str string
								switch meta.Value.(type) {
								case string:
									str = fmt.Sprintf("\"%v\"", meta.Value)
								case int, float32, float64, bool:
									switch meta.Value.(type) {
									case bool:
										str = fmt.Sprintf("%v", meta.Value.(bool))
									default:
										str = fmt.Sprintf("%d", meta.Value)
									}
								}
								p.DefaultValue = &str
							case "array":
								if v, ok := meta.Value.(bool); ok {
									p.IsArray = v
								}
							case "optional":
								if v, ok := meta.Value.(bool); ok {
									p.IsRequired = !v
								}
							case "type":
								if v, ok := meta.Value.(string); ok {
									p.Type = v
								} else {
									return nil, fmt.Errorf("prop \"%v\" type value at model \"%v\" must be a string", p.PropName, m.ModelName)
								}
							}
						}
					}
					m.Props = append(m.Props, p)
				}
			}
			root.Models[key] = m
		}
	}

	return root, nil
}
