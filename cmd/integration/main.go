package main

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/log_analyzer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", os.Args[0]).Logger()

	la := log_analyzer.NewLA()

	cmds := [][]string{
		[]string{"sync", "./bin/game_sync", "./configs/config_sync.json"},
		[]string{"core", "./bin/game_core", "./configs/config_core.json"},
		[]string{"api", "./bin/game_api", "./configs/config_api.json"},
		[]string{"auth", "./bin/game_auth", "./configs/config_auth.json"},
		[]string{"revoker", "./bin/game_revoker", "./configs/config_revoker.json"},
		[]string{"tool", "./bin/game_tool", "./configs/config_tool.json"},
	}

	defer la.Close()
	for _, args := range cmds {
		if err := la.NewProcess(args[0], args[1:]...); err != nil {
			log.Error().Err(err).Msg("failed to start")
			return
		}
	}

	laClient := log_analyzer.NewLA()
	defer laClient.Close()
	if err := laClient.NewProcess(
		"client",
		"./bin/game_client",
		"./configs/config_client.json",
	); err != nil {
		log.Error().Err(err).Msg("failed to start")
		return
	}

	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec

	log.Info().Msg("integration up")

	authService := auth.NewService("https://localhost:8080")

	la.Consume(7)

	if err := Case0(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	if err := Case1(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
}
