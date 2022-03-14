package ast

import (
	"strings"
)

type Options map[string]string

type Node interface {
	Name() string
	Options() Options
}

type RootNode struct {
	Models      ModelMap
	RootOptions Options
}

func (m *RootNode) Name() string     { return "__RootNode__" }
func (m *RootNode) Options() Options { return m.RootOptions }

type Model struct {
	ModelName    string
	Props        []*ModelProp
	ModelOptions Options
}

func (m *Model) Name() string     { return m.ModelName }
func (m *Model) Options() Options { return m.ModelOptions }

func (m *Model) GetProp(propName string) *ModelProp {
	for _, p := range m.Props {
		if p.PropName == propName {
			return p
		}
	}
	return nil
}

type ModelProp struct {
	PropName     string
	IsRequired   bool
	IsArray      bool
	DefaultValue *string
	Type         string
	PropOptions  Options
}

func (m *ModelProp) Name() string     { return m.PropName }
func (m *ModelProp) Options() Options { return m.PropOptions }

type ModelMap map[string]*Model

func (models ModelMap) ModelDependencies(model interface{}) []*Model {
	deps := make([]*Model, 0)

	var modelName string
	if m, ok := model.(*Model); ok {
		modelName = m.ModelName
	} else if m, ok := model.(string); ok {
		modelName = m
	} else {
		return deps
	}

	m, ok := models[modelName]
	if !ok {
		return make([]*Model, 0)
	}

	for _, p := range m.Props {
		if mm, ok := models[string(p.Type)]; ok {
			ok := false
			for _, d := range deps {
				if d.ModelName == mm.ModelName {
					ok = true
					break
				}
			}
			if !ok {
				deps = append(deps, mm)
			}
		}
	}
	return deps
}

func GetNodeOption(r Node, optionName string) string {
	o, ok := r.Options()[optionName]
	if !ok {
		return ""
	}
	return strings.TrimPrefix(strings.TrimSuffix(o, "\""), "\"")
}
