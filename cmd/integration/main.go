package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	gulid "github.com/elojah/game_01/pkg/ulid"
)

var (
	me = connectInfo{
		Username: "ITEST_me_username",
		Password: "ITEST_me_password",
		PCName:   "ITEST_me_pcname",
		PCType:   gulid.MustParse("01CE3J5M6QMP5A4C0GTTYFYANP"),
	}
	opponent0 = connectInfo{
		Username: "ITEST_o0_username",
		Password: "ITEST_o0_password",
		PCName:   "ITEST_o0_pcname",
		PCType:   gulid.MustParse("01CE3J5M6QMP5A4C0GTTYFYANP"),
	}
)

func main() {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", os.Args[0]).Logger()

	la := NewLogAnalyzer()

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

	laClient := NewLogAnalyzer()
	defer laClient.Close()
	if err := laClient.NewProcess(
		"client",
		"./bin/game_client",
		"./configs/config_client.json",
	); err != nil {
		log.Error().Err(err).Msg("failed to start")
		return
	}

	log.Info().Msg("integration up")

	if err := expectUp(la); err != nil {
		log.Error().Err(err).Msg("up")
		return
	}
	if err := expectUpClient(laClient); err != nil {
		log.Error().Err(err).Msg("client up")
		return
	}
	log.Info().Msg("up ok")

	if err := expectStatic(la); err != nil {
		log.Error().Err(err).Msg("static")
		return
	}
	log.Info().Msg("static ok")

	/*
		#Create and use entity
		- Create account
		- Create new pj
		- Connect to PJ
		- Check client feedback
	*/

	tok, ent, err := expectConnect(la, me)
	if err != nil {
		log.Error().Err(err).Msg("connect")
		return
	}
	log.Info().Msg("connect ok")

	entClient, err := expectState(laClient, ent)
	if err != nil {
		log.Error().Err(err).Msg("state")
		return
	}
	log.Info().Msg("state ok")

	time.Sleep(1 * time.Millisecond)

	/*
		#Test entity move
		- FAIL Move same sector too far
		- SUCCESS Move same sector
		- FAIL Move not neighbour sector
		- SUCCESS Move with tool (no check)
		- FAIL Move neighbour sector too far
		- SUCCESS Move neighbour sector
	*/

	if err := expectMoveSameSectorTooFar(la, laClient, tok, entClient); err != nil {
		log.Error().Err(err).Msg("move same sector too far")
		return
	}
	log.Info().Msg("move same sector too far ok")

	if entClient, err = expectMoveSameSector(la, laClient, tok, entClient); err != nil {
		log.Error().Err(err).Msg("move same sector")
		return
	}
	log.Info().Msg("move same sector ok")

	if err := expectMoveNotNeighbourSector(la, laClient, tok, entClient); err != nil {
		log.Error().Err(err).Msg("move not neighbour sector")
		return
	}
	log.Info().Msg("move not neighbour sector ok")

	if entClient, err = expectToolEntityMove(la, entClient); err != nil {
		log.Error().Err(err).Msg("move entity with tool")
		return
	}
	log.Info().Msg("move entity with tool ok")

	if err := expectMoveNeighbourTooFar(la, laClient, tok, entClient); err != nil {
		log.Error().Err(err).Msg("move neighbour sector too far")
		return
	}
	log.Info().Msg("move neighbour sector too far ok")

	if entClient, err = expectMoveNeighbourSector(la, laClient, tok, entClient); err != nil {
		log.Error().Err(err).Msg("move neighbour sector")
		return
	}
	log.Info().Msg("move neighbour sector ok")

	/*
		#Test cast mechanism
		- Create an opponent
	*/

	toko0, ento0, err := expectConnect(la, opponent0)
	if err != nil {
		log.Error().Err(err).Msg("connect")
		return
	}
	log.Info().Msg("connect ok")
	_, _ = toko0, ento0

	_ = entClient
	// if err := expectDisconnect(la, tok); err != nil {
	// 	log.Error().Err(err).Msg("disconnect")
	// 	return
	// }
	// log.Info().Msg("disconnect ok")
	log.Info().Msg("integration ok")
}
