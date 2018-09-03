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

	defer a.Close()
	for _, args := range cmds {
		if err := a.Cmd(args...); err != nil {
			log.Error().Err(err).Msg("failed to start")
			return
		}
	}

	log.Info().Msg("integration up")

	if err := expectUp(a); err != nil {
		log.Error().Err(err).Msg("up")
		return
	}
	log.Info().Msg("up ok")

	if err := expectTool(a); err != nil {
		log.Error().Err(err).Msg("tool")
		return
	}
	log.Info().Msg("tool ok")

	tok, ent, err := expectAuthUp(a)
	if err != nil {
		log.Error().Err(err).Msg("auth up")
		return
	}
	log.Info().Msg("auth up ok")

	if err := expectAPI(a, tok, ent); err != nil {
		log.Error().Err(err).Msg("api")
		return
	}
	log.Info().Msg("api ok")

	if err := expectAuthDown(a, tok); err != nil {
		log.Error().Err(err).Msg("auth down")
		return
	}
	log.Info().Msg("auth down ok")

	log.Info().Msg("integration ok")
}
