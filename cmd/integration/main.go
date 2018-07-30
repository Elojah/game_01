package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", os.Args[0]).Logger()

	a := NewLogAnalyzer()

	cmds := [][]string{
		[]string{"./bin/game_sync", "./configs/config_sync.json"},
		[]string{"./bin/game_core", "./configs/config_core.json"},
		[]string{"./bin/game_api", "./configs/config_api.json"},
		[]string{"./bin/game_auth", "./configs/config_auth.json"},
		[]string{"./bin/game_revoker", "./configs/config_revoker.json"},
		[]string{"./bin/game_tool", "./configs/config_tool.json"},
	}

	for _, args := range cmds {
		if err := a.Cmd(args...); err != nil {
			log.Error().Err(err).Msg("failed to start")
			return
		}
	}

	log.Info().Msg("integration up")

	select {}
}
