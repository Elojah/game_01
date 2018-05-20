package main

import (
	"encoding/json"
	"fmt"
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
func (t *template) init() error {

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
		return err
	}
	return nil
}

func (t *template) ShowTemplates(cmd *cobra.Command, args []string) {

	if err := t.init(); err != nil {
		return
	}

	for _, arg := range args {

		if arg == "skills" {
			skills, err := t.ListSkillTemplate()
			if err != nil {
				t.logger.Error().Err(err).Msg("failed to retrieve skills")
				return
			}
			t.logger.Info().Int("skills", len(skills)).Msg("found")
			for _, skill := range skills {
				raw, err := json.Marshal(skill)
				if err != nil {
					t.logger.Error().Err(err).Msg("failed to retrieve skills")
					return
				}
				fmt.Println(string(raw))
			}
		}

		if arg == "entities" {
			entities, err := t.ListEntityTemplate()
			if err != nil {
				t.logger.Error().Err(err).Msg("failed to retrieve entities")
				return
			}
			t.logger.Info().Int("entities", len(entities)).Msg("found")
			for _, entity := range entities {
				raw, err := json.Marshal(entity)
				if err != nil {
					t.logger.Error().Err(err).Msg("failed to retrieve entities")
					return
				}
				fmt.Println(string(raw))
			}
		}

	}

	t.logger.Info().Msg("done")
}

func (t *template) AddTemplates(cmd *cobra.Command, args []string) {

	if err := t.init(); err != nil {
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
