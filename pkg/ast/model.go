package ast

import (
	"fmt"
	"strings"

	"github.com/juanvillacortac/rosetta/pkg/parser"
	"github.com/juanvillacortac/rosetta/pkg/types"

	proto "github.com/yoheimuta/go-protoparser/v4/parser"
)

type Options map[string]string

type Node interface {
	Name() string
	Options() Options
}

type Model struct {
	ModelName    string
	Props        []ModelProp
	ModelOptions Options
}

func (m *Model) Name() string     { return m.ModelName }
func (m *Model) Options() Options { return m.ModelOptions }

type ModelProp struct {
	PropName     string
	IsRequired   bool
	IsArray      bool
	DefaultValue *string
	Type         string
	PropOptions  map[string]string
}

func (m *ModelProp) Name() string     { return m.PropName }
func (m *ModelProp) Options() Options { return m.PropOptions }

type ModelMap map[string]*Model

func (models ModelMap) GetModelDeps(modelName string, deps []string) []string {
	if deps == nil {
		deps = make([]string, 0)
	}
	clone := make([]string, len(deps))
	copy(clone, deps)

	m, ok := models[modelName]
	if !ok {
		return nil
	}

	for _, p := range m.Props {
		if mm, ok := models[string(p.Type)]; ok {
			clone = append(clone, mm.ModelName)
		}
	}
	return clone
}

func GetModelsFromProto(program *proto.Proto) (ModelMap, error) {
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
			var def *string
			options := make(Options)
			for _, o := range f.FieldOptions {
				if o.OptionName == "default" {
					def = &o.Constant
				}
				if strings.HasPrefix(o.OptionName, "(") && strings.HasSuffix(o.OptionName, ")") {
					name := strings.TrimPrefix(strings.TrimSuffix(o.OptionName, ")"), "(")
					options[name] = o.Constant
				}
			}
			prop := &ModelProp{
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
			props = append(props, *prop)
		}

		if _, exist := nodes[m.MessageName]; exist {
			err := fmt.Errorf("message \"%v\" is already defined", m.MessageName)
			return nil, err
		}

		options := make(Options)
		for _, b := range m.MessageBody {
			if o, ok := b.(*proto.Option); ok {
				if strings.HasPrefix(o.OptionName, "(") && strings.HasSuffix(o.OptionName, ")") {
					name := strings.TrimPrefix(strings.TrimSuffix(o.OptionName, ")"), "(")
					options[name] = o.Constant
				}
			}
		}

		nodes[m.MessageName] = &Model{
			ModelName:    m.MessageName,
			Props:        props,
			ModelOptions: options,
		}
	}
	return nodes, nil
}
