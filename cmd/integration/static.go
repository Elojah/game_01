package main

import (
	"github.com/elojah/game_01/cmd/integration/tool"
	"github.com/pkg/errors"
)

// Static creates static prerequired data.
func Static(s *tool.Service) error {
	if err := s.AddEntityTemplates("./static/entity_templates.json"); err != nil {
		return errors.Wrap(err, "static")
	}
	if err := s.AddAbilityTemplates("./static/ability_templates.json"); err != nil {
		return errors.Wrap(err, "static")
	}
	if err := s.AddAbilityStarter("./static/ability_starter.json"); err != nil {
		return errors.Wrap(err, "static")
	}
	if err := s.AddSector("./static/sector.json"); err != nil {
		return errors.Wrap(err, "static")
	}
	if err := s.AddSpawn("./static/spawn.json"); err != nil {
		return errors.Wrap(err, "static")
	}
	return nil
}
