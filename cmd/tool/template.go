package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/elojah/game_01"
	redisx "github.com/elojah/game_01/storage/redis"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

type template struct {
	game.EntityTemplateService
	game.SkillTemplateService

	config   string
	entities string
	skills   string

	logger zerolog.Logger
}

// run template tool.
func (t *template) run(cmd *cobra.Command, args []string) {

	t.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	launchers := services.Launchers{}

	// redis
	rd := redis.Service{}
	rdl := rd.NewLauncher(redis.Namespaces{
		Redis: "redis",
	}, "redis")
	launchers = append(launchers, rdl)
	rdx := redisx.NewService(&rd)

	t.EntityTemplateService = rdx
	t.SkillTemplateService = rdx

	if err := launchers.Up(t.config); err != nil {
		t.logger.Error().Err(err).Str("filename", t.config).Msg("failed to start")
		return
	}

	if t.entities != "" {
		t.CreateEntities()
	}
	if t.skills != "" {
		t.CreateSkills()
	}

	t.logger.Info().Msg("done")
}

func (t *template) CreateEntities() {

	raw, err := ioutil.ReadFile(t.entities)
	if err != nil {
		t.logger.Error().Err(err).Str("entities", t.entities).Msg("failed to read entities file")
		return
	}
	var entities []game.EntityTemplate
	if err := json.Unmarshal(raw, &entities); err != nil {
		t.logger.Error().Err(err).Str("entities", t.entities).Msg("invalid JSON")
		return
	}

	t.logger.Info().Int("entities", len(entities)).Msg("found")

	for _, tpl := range entities {
		if err := t.SetEntityTemplate(tpl); err != nil {
			t.logger.Error().Err(err).Str("template", tpl.ID.String()).Msg("failed to set template")
			return
		}
	}
}

func (t *template) CreateSkills() {

	raw, err := ioutil.ReadFile(t.skills)
	if err != nil {
		t.logger.Error().Err(err).Str("skills", t.skills).Msg("failed to read skills file")
		return
	}
	var skills []game.SkillTemplate
	if err := json.Unmarshal(raw, &skills); err != nil {
		t.logger.Error().Err(err).Str("skills", t.skills).Msg("invalid JSON")
		return
	}

	t.logger.Info().Int("skills", len(skills)).Msg("found")

	for _, tpl := range skills {
		if err := t.SetSkillTemplate(tpl); err != nil {
			t.logger.Error().Err(err).Str("template", tpl.ID.String()).Msg("failed to set template")
			return
		}
	}
}
