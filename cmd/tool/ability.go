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

type abilityWithEntity struct {
	game.Ability
	EntityID game.ID `json:"entity_id"`
}

type ability struct {
	game.AbilityMapper

	config    string
	abilities string

	logger zerolog.Logger
}

// run ability tool.
func (a *ability) init() error {

	a.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	launchers := services.Launchers{}

	// redis
	rd := redis.Service{}
	rdl := rd.NewLauncher(redis.Namespaces{
		Redis: "redis",
	}, "redis")
	launchers = append(launchers, rdl)
	rdx := redisx.NewService(&rd)

	a.AbilityMapper = rdx

	if err := launchers.Up(a.config); err != nil {
		a.logger.Error().Err(err).Str("filename", a.config).Msg("failed to start")
		return err
	}
	return nil
}

func (a *ability) AddAbilities(cmd *cobra.Command, args []string) {

	if err := a.init(); err != nil {
		return
	}

	a.CreateAbilities()

	a.logger.Info().Msg("done")
}

func (a *ability) CreateAbilities() {

	raw, err := ioutil.ReadFile(a.abilities)
	if err != nil {
		a.logger.Error().Err(err).Str("abilities", a.abilities).Msg("failed to read abilities file")
		return
	}
	var abilities []abilityWithEntity
	if err := json.Unmarshal(raw, &abilities); err != nil {
		a.logger.Error().Err(err).Str("abilities", a.abilities).Msg("invalid JSON")
		return
	}

	a.logger.Info().Int("abilities", len(abilities)).Msg("found")

	for _, sk := range abilities {
		if err := a.SetAbility(sk.Ability, sk.EntityID); err != nil {
			a.logger.Error().Err(err).Str("ability", sk.ID.String()).Msg("failed to set ability")
			return
		}
	}
}
