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
	game.EntityTemplateMapper
	game.AbilityTemplateMapper

	config    string
	entities  string
	abilities string

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

	t.EntityTemplateMapper = rdx
	t.AbilityTemplateMapper = rdx

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

		if arg == "abilities" {
			abilities, err := t.ListAbilityTemplate()
			if err != nil {
				t.logger.Error().Err(err).Msg("failed to retrieve abilities")
				return
			}
			t.logger.Info().Int("abilities", len(abilities)).Msg("found")
			for _, ability := range abilities {
				raw, err := json.Marshal(ability)
				if err != nil {
					t.logger.Error().Err(err).Msg("failed to retrieve abilities")
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
	if t.abilities != "" {
		t.CreateAbilities()
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

func (t *template) CreateAbilities() {

	raw, err := ioutil.ReadFile(t.abilities)
	if err != nil {
		t.logger.Error().Err(err).Str("abilities", t.abilities).Msg("failed to read abilities file")
		return
	}
	var abilities []game.AbilityTemplate
	if err := json.Unmarshal(raw, &abilities); err != nil {
		t.logger.Error().Err(err).Str("abilities", t.abilities).Msg("invalid JSON")
		return
	}

	t.logger.Info().Int("abilities", len(abilities)).Msg("found")

	for _, tpl := range abilities {
		if err := t.SetAbilityTemplate(tpl); err != nil {
			t.logger.Error().Err(err).Str("template", tpl.ID.String()).Msg("failed to set template")
			return
		}
	}
}
