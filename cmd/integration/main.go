package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/client"
	"github.com/elojah/game_01/cmd/integration/log_analyzer"
	"github.com/elojah/game_01/cmd/integration/tool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", os.Args[0]).Logger()

	la := loganalyzer.NewLA()

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

	authService := auth.NewService("https://localhost:8080")
	if err := Case0(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	if err := Case1(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	if err := Case2(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	if err := Case3(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	if err := Case4(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}

	clientService := client.NewService(laClient)
	if err := Case5(authService, clientService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	if err := Case6(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	if err := Case7(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
}
