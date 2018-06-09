package main

import (
	"os"

	"github.com/rs/zerolog"

	game "github.com/elojah/game_01"
	redisx "github.com/elojah/game_01/storage/redis"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

type sector struct {
	game.SectorMapper

	config  string
	sectors string

	logger zerolog.Logger
}

// run sector tool.
func (s *sector) init() error {

	s.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	launchers := services.Launchers{}

	// redis
	rd := redis.Service{}
	rdl := rd.NewLauncher(redis.Namespaces{
		Redis: "redis",
	}, "redis")
	launchers = append(launchers, rdl)
	rdx := redisx.NewService(&rd)

	s.SectorMapper = rdx

	if err := launchers.Up(s.config); err != nil {
		s.logger.Error().Err(err).Str("filename", s.config).Msg("failed to start")
		return err
	}
	return nil
}
