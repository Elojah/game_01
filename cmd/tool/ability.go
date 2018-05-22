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

	config   string
	abilitys string

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

func (a *ability) AddAbilitys(cmd *cobra.Command, args []string) {

	if err := a.init(); err != nil {
		return
	}

	a.CreateAbilitys()

	a.logger.Info().Msg("done")
}

func (a *ability) CreateAbilitys() {

	raw, err := ioutil.ReadFile(a.abilitys)
	if err != nil {
		a.logger.Error().Err(err).Str("abilitys", a.abilitys).Msg("failed to read abilitys file")
		return
	}
	var abilitys []abilityWithEntity
	if err := json.Unmarshal(raw, &abilitys); err != nil {
		a.logger.Error().Err(err).Str("abilitys", a.abilitys).Msg("invalid JSON")
		return
	}

	a.logger.Info().Int("abilitys", len(abilitys)).Msg("found")

	for _, sk := range abilitys {
		if err := a.SetAbility(sk.Ability, sk.EntityID); err != nil {
			a.logger.Error().Err(err).Str("ability", sk.ID.String()).Msg("failed to set ability")
			return
		}
	}
}
