package game

import (
	"encoding/json"
)

// AbilityTemplate alias an entity.
// It represents semi static data. When creating abilities, those templates are used.
type AbilityTemplate Ability

// AbilityTemplateMapper is an interface for AbilityTemplate object.
type AbilityTemplateMapper interface {
	SetAbilityTemplate(AbilityTemplate) error
	GetAbilityTemplate(AbilityTemplateSubset) (AbilityTemplate, error)
	ListAbilityTemplate() ([]AbilityTemplate, error)
}

// AbilityTemplateSubset is a subset to retrieve one template.
type AbilityTemplateSubset struct {
	Type AbilityType
}

// UnmarshalJSON adds a new field for components `struct` to determine component type.
func (a *AbilityTemplate) UnmarshalJSON(raw []byte) error {
	var alias struct {
		Ability
		Components []json.RawMessage `json:"components"`
	}
	if err := json.Unmarshal(raw, &alias); err != nil {
		return err
	}

	*a = AbilityTemplate(alias.Ability)
	a.Components = make([]AbilityComponent, len(alias.Components))

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
