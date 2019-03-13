package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/cmd/integration/loganalyzer"
	"github.com/elojah/game_01/cmd/integration/tool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", os.Args[0]).Logger()

	la := loganalyzer.NewLA()

	cmds := [][]string{
		{"core", "./bin/game_core", "./configs/config_core.json"},
		{"api", "./bin/game_api", "./configs/config_api.json"},
		{"auth", "./bin/game_auth", "./configs/config_auth.json"},
		{"revoker", "./bin/game_revoker", "./configs/config_revoker.json"},
		{"tool", "./bin/game_tool", "./configs/config_tool.json"},
	}

	defer la.Close()
	for _, args := range cmds {
		if err := la.NewProcess(args[0], args[1:]...); err != nil {
			log.Error().Err(err).Msg("failed to start")
			return
		}
	}

	laSync := loganalyzer.NewLA()
	defer laSync.Close()
	if err := laSync.NewProcess(
		"sync",
		"./bin/game_sync",
		"./configs/config_sync.json",
	); err != nil {
		log.Error().Err(err).Msg("failed to start")
		return
	}

	laClient := loganalyzer.NewLA()
	defer laClient.Close()
	if err := laClient.NewProcess(
		"client",
		"./bin/game_client",
		"./configs/config_client.json",
	); err != nil {
		log.Error().Err(err).Msg("failed to start")
		return
	}

	// wait for processes to listen
	time.Sleep(500 * time.Millisecond)

	// ignore certificate validity
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec

	log.Info().Msg("integration up")

	toolService := tool.NewService("https://localhost:8081")
	if err := Static(toolService); err != nil {
		log.Error().Err(err).Msg("static failure")
		return
	}
	log.Info().Msg("static ok")

	authService := auth.NewService("https://localhost:8080")
	if err := Case0(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case0 ok")
	if err := Case1(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case1 ok")
	if err := Case2(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case2 ok")
	if err := Case3(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case3 ok")
	if err := Case4(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case4 ok")
	// TODO reorder
	if err := Case9(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case9 ok")

	clientService := client.NewService(laClient)
	if err := Case5(authService, clientService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case5 ok")
	if err := Case6(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case6 ok")
	if err := Case7(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case7 ok")
	if err := Case8(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case8 ok")
	log.Info().Msg("integration ok")
}
