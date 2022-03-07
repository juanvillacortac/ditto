package proto

import (
	"fmt"
	"strings"

	"github.com/juanvillacortac/rosetta/pkg/ast"
	"github.com/juanvillacortac/rosetta/pkg/types"
	"github.com/juanvillacortac/rosetta/pkg/utils"
	protoparser "github.com/yoheimuta/go-protoparser/v4/parser"
)

func GetRootNodeFromProto(program *protoparser.Proto) (*ast.RootNode, error) {
	return GetRootNode(program)
}

func GetRootNode(program *protoparser.Proto) (*ast.RootNode, error) {
	root := &ast.RootNode{
		RootOptions: make(ast.Options),
	}
	models := make(ast.ModelMap)
	messages := GetMessages(program.ProtoBody)
	options := GetOptions(program.ProtoBody)
	for _, o := range options {
		if name, wrapped := utils.UnwrapString(o.OptionName, "(", ")"); wrapped {
			root.RootOptions[name] = o.Constant
		}
	}
	for _, m := range messages {
		props := make([]*ast.ModelProp, 0)
		for _, mb := range m.MessageBody {
			switch mb.(type) {
			case *protoparser.Field:
				f := mb.(*protoparser.Field)
				var def *string
				options := make(ast.Options)
				for _, o := range f.FieldOptions {
					if o.OptionName == "default" {
						def = &o.Constant
					}
					if strings.HasPrefix(o.OptionName, "(") && strings.HasSuffix(o.OptionName, ")") {
						name := strings.TrimPrefix(strings.TrimSuffix(o.OptionName, ")"), "(")
						options[name] = o.Constant
					}
				}
				prop := &ast.ModelProp{
					PropName:     f.FieldName,
					IsRequired:   f.IsRequired,
					IsArray:      f.IsRepeated,
					Type:         f.Type,
					DefaultValue: def,
					PropOptions:  options,
				}

				if fType, ok := types.TypesMap[f.Type]; ok {
					prop.Type = string(fType)
				}
				props = append(props, prop)
			}
		}

		if _, exist := models[m.MessageName]; exist {
			err := fmt.Errorf("message \"%v\" is already defined", m.MessageName)
			return nil, err
		}

		options := make(ast.Options)
		for _, b := range m.MessageBody {
			if o, ok := b.(*protoparser.Option); ok {
				if strings.HasPrefix(o.OptionName, "(") && strings.HasSuffix(o.OptionName, ")") {
					name := strings.TrimPrefix(strings.TrimSuffix(o.OptionName, ")"), "(")
					options[name] = o.Constant
				}
			}
		}

		models[m.MessageName] = &ast.Model{
			ModelName:    m.MessageName,
			Props:        props,
			ModelOptions: options,
		}
	}
	root.Models = models
	return root, nil
}
