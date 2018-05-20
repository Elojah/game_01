package storage

import (
	"github.com/elojah/game_01"
)

// NewSkill convert a game.Skill into a storage Skill.
func NewSkill(a game.Skill) *Skill {
	return &Skill{
		ID:            [16]byte(a.ID),
		Type:          [16]byte(a.Type),
		Name:          a.Name,
		MPConsumption: a.MPConsumption,
		DirectDamage:  a.DirectDamage,
		DirectHeal:    a.DirectHeal,
		CD:            a.CD,
		CurrentCD:     a.CurrentCD,
	}
}

// Domain converts a storage Skill into a game Skill.
func (a Skill) Domain() game.Skill {
	return game.Skill{
		ID:            game.ID(a.ID),
		Type:          game.SkillType(a.Type),
		Name:          a.Name,
		MPConsumption: a.MPConsumption,
		DirectDamage:  a.DirectDamage,
		DirectHeal:    a.DirectHeal,
		CD:            a.CD,
		CurrentCD:     a.CurrentCD,
	}
}
