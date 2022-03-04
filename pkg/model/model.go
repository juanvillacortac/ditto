package model

import (
	"fmt"

	"github.com/juanvillacortac/rosetta/pkg/parser"
	"github.com/juanvillacortac/rosetta/pkg/types"

	proto "github.com/yoheimuta/go-protoparser/v4/parser"
)

type Model struct {
	Name  string
	Props []ModelProp
}

type ModelProp struct {
	Name       string
	IsRequired bool
	IsArray    bool
	Type       types.Type
}

type ModelMap map[string]Model

func GetModelsFromPrograms(program *proto.Proto) (ModelMap, error) {
	messages := parser.GetMessages(program.ProtoBody)
	return GetModels(messages)
}

func GetModels(messages []*proto.Message) (ModelMap, error) {
	nodes := ModelMap{}
	for _, m := range messages {
		props := make([]ModelProp, 0)

		for _, mb := range m.MessageBody {
			f, ok := mb.(*proto.Field)
			if !ok {
				continue
			}
			prop := &ModelProp{
				Name:       f.FieldName,
				IsRequired: f.IsRequired,
				IsArray:    f.IsRepeated,
				Type:       f.Type,
			}

			if fType, ok := types.TypesMap[f.Type]; ok {
				prop.Type = fType
			}
			props = append(props, *prop)
		}

		if _, exist := nodes[m.MessageName]; exist {
			err := fmt.Errorf("Message \"%v\" is already defined", m.MessageName)
			return nil, err
		}

		nodes[m.MessageName] = Model{
			Name:  m.MessageName,
			Props: props,
		}
	}
	return nodes, nil
}
