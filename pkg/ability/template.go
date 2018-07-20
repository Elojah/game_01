package ability

import (
	"encoding/json"
)

// Template alias an entity.
// It represents semi static data. When creating abilities, those templates are used.
type Template = A

// TemplateMapper is an interface for Template object.
type TemplateMapper interface {
	SetAbilityTemplate(Template) error
	GetAbilityTemplate(TemplateSubset) (Template, error)
	ListAbilityTemplate() ([]Template, error)
}

// TemplateSubset is a subset to retrieve one template.
type TemplateSubset struct {
	Type Type
}

// UnmarshalJSON adds a new field for components `struct` to determine component type.
func (a *Template) UnmarshalJSON(raw []byte) error {
	var alias struct {
		A
		Components []json.RawMessage `json:"components"`
	}
	if err := json.Unmarshal(raw, &alias); err != nil {
		return err
	}

	*a = Template(alias.A)
	a.Components = make([]Component, len(alias.Components))

	for i, component := range alias.Components {
		var componentStruct struct {
			Structure string `json:"struct"`
		}
		if err := json.Unmarshal([]byte(component), &componentStruct); err != nil {
			return err
		}
		switch componentStruct.Structure {
		case "HealDirect":
			var s HealDirect
			if err := json.Unmarshal([]byte(component), &s); err != nil {
				return err
			}
			a.Components[i] = s
		case "DamageDirect":
			var s DamageDirect
			if err := json.Unmarshal([]byte(component), &s); err != nil {
				return err
			}
			a.Components[i] = s
		case "HealOverTime":
			var s HealOverTime
			if err := json.Unmarshal([]byte(component), &s); err != nil {
				return err
			}
			a.Components[i] = s
		case "DamageOverTime":
			var s DamageOverTime
			if err := json.Unmarshal([]byte(component), &s); err != nil {
				return err
			}
			a.Components[i] = s
		default:
			return &json.UnsupportedValueError{Str: "struct"}
		}
	}
	return nil
}
