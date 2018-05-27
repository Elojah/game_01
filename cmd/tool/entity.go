package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/elojah/game_01"
	tile38x "github.com/elojah/game_01/storage/tile38"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

type entity struct {
	game.EntityPositionMapper

	config    string
	entities  string
	positions string
	radius    float64

	logger zerolog.Logger
}

// run entity tool.
func (e *entity) init() error {

	e.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	launchers := services.Launchers{}

	// redis
	t38 := redis.Service{}
	t38l := t38.NewLauncher(redis.Namespaces{
		Redis: "tile38",
	}, "tile38")
	launchers = append(launchers, t38l)
	t38x := tile38x.NewService(&t38)

	e.EntityPositionMapper = t38x

	if err := launchers.Up(e.config); err != nil {
		e.logger.Error().Err(err).Str("filename", e.config).Msg("failed to start")
		return err
	}
	return nil
}

func (e *entity) Start(cmd *cobra.Command, args []string) {
	if err := e.init(); err != nil {
		return
	}

	if e.entities != "" {
		e.SpawnEntities()
	}
	if e.positions != "" {
		e.ShowEntities()
	}

	e.logger.Info().Msg("done")

}

func (e *entity) SpawnEntities() {
	raw, err := ioutil.ReadFile(e.entities)
	if err != nil {
		e.logger.Error().Err(err).Str("entities", e.entities).Msg("failed to read entities file")
		return
	}
	var entities []game.Entity
	if err := json.Unmarshal(raw, &entities); err != nil {
		e.logger.Error().Err(err).Str("entities", e.entities).Msg("invalid JSON")
		return
	}

	e.logger.Info().Int("entities", len(entities)).Msg("found")

	for _, en := range entities {
		if err := e.SetEntityPosition(en, 0); err != nil {
			e.logger.Error().Err(err).Str("entity", en.ID.String()).Msg("failed to set entity")
			return
		}
	}

}

func (e *entity) ShowEntities() {
	raw, err := ioutil.ReadFile(e.positions)
	if err != nil {
		e.logger.Error().Err(err).Str("positions", e.positions).Msg("failed to read positions file")
		return
	}
	var positions []game.Vec3
	if err := json.Unmarshal(raw, &positions); err != nil {
		e.logger.Error().Err(err).Str("positions", e.positions).Msg("invalid JSON")
		return
	}

	e.logger.Info().Int("positions", len(positions)).Msg("found")

	for _, p := range positions {
		entities, err := e.ListEntityPosition(game.EntityPositionSubset{
			Position: p,
			Radius:   e.radius,
		})
		if err != nil {
			e.logger.Error().Err(err).Msg("failed to retrieve entities")
			return
		}
		e.logger.Info().Int("entities", len(entities)).Msg("found")
		for _, en := range entities {
			raw, err := json.Marshal(en)
			if err != nil {
				e.logger.Error().Err(err).Msg("failed to marshal entities")
				return
			}
			fmt.Println(string(raw))
		}
	}
}
