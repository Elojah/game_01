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

	config   string
	entities string
	abilitys string

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

		if arg == "abilitys" {
			abilitys, err := t.ListAbilityTemplate()
			if err != nil {
				t.logger.Error().Err(err).Msg("failed to retrieve abilitys")
				return
			}
			t.logger.Info().Int("abilitys", len(abilitys)).Msg("found")
			for _, ability := range abilitys {
				raw, err := json.Marshal(ability)
				if err != nil {
					t.logger.Error().Err(err).Msg("failed to retrieve abilitys")
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
	if t.abilitys != "" {
		t.CreateAbilitys()
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

func (t *template) CreateAbilitys() {

	raw, err := ioutil.ReadFile(t.abilitys)
	if err != nil {
		t.logger.Error().Err(err).Str("abilitys", t.abilitys).Msg("failed to read abilitys file")
		return
	}
	var abilitys []game.AbilityTemplate
	if err := json.Unmarshal(raw, &abilitys); err != nil {
		t.logger.Error().Err(err).Str("abilitys", t.abilitys).Msg("invalid JSON")
		return
	}

	t.logger.Info().Int("abilitys", len(abilitys)).Msg("found")

	for _, tpl := range abilitys {
		if err := t.SetAbilityTemplate(tpl); err != nil {
			t.logger.Error().Err(err).Str("template", tpl.ID.String()).Msg("failed to set template")
			return
		}
	}
}
