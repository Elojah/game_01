package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/elojah/game_01/cmd/integration/auth"
	"github.com/elojah/game_01/cmd/integration/cases"
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
	if err := cases.Subscribe(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case subscribe ok")
	if err := cases.Sign(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case sign ok")
	if err := cases.SignBis(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case sign_bis ok")
	if err := cases.CreatePC(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case create_pc ok")
	if err := cases.ConnectPC(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case connect_pc ok")
	if err := cases.DelPC(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case del_pc ok")
	if err := cases.DelPCBis(authService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case del_pc_bis ok")

	clientService := client.NewService(laClient)
	if err := cases.Move(authService, clientService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case move ok")
	if err := cases.MoveSector(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case move_sector ok")
	if err := cases.Cast(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case cast ok")
	if err := cases.CastSector(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case cast_sector ok")
	if err := cases.Loot(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case loot ok")
	if err := cases.Consume(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case consume ok")
	if err := cases.ConsumeCast(authService, clientService, toolService); err != nil {
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case consume_cast ok")
	if err := cases.Spawn(authService, clientService, toolService); err != nil {
		laClient.Consume(500)
		log.Error().Err(err).Msg("case failure")
		return
	}
	log.Info().Msg("case spawn ok")

	log.Info().Msg("integration ok")
}
