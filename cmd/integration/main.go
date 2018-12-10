package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/ability"
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
		log.Error().Err(err).Msg("connect o0")
		return
	}
	log.Info().Msg("connect o0 ok")
	if ento0, err = expectToolo0Move(la, ento0); err != nil {
		log.Error().Err(err).Msg("move o0 with tool")
		return
	}
	log.Info().Msg("move o0 with tool ok")
	ab := ability.A{
		ID:                gulid.NewID(),
		Type:              gulid.MustParse("01CP2Z4SDEWZK8YF29E07GPDVC"),
		Animation:         gulid.MustParse("00000000000000000000000000"),
		Name:              "int_test_ability",
		MPConsumption:     5,
		PostMPConsumption: 0,
		CD:                2000,
		LastUsed:          0,
		CastTime:          1000,
		Components: map[string]ability.Component{
			"01CYBX32YGJ4A4T4SAMMQKQS1H": ability.Component{
				Effects: []ability.Effect{
					ability.Effect{
						Damage: &ability.Damage{
							Element: ability.Time,
							Amount:  8,
						},
					},
				},
				NTargets:   1,
				Range:      50,
				NPositions: 0,
			},
		},
	}
	if err := expectToolSetAbility(la, ab, ent); err != nil {
		log.Error().Err(err).Msg("set ability with tool")
		return
	}
	log.Info().Msg("set ability with tool ok")
	if entClient, err = expectCast(la, laClient, tok, ab.ID, entClient, ento0); err != nil {
		log.Error().Err(err).Msg("cast on o0")
		return
	}
	log.Info().Msg("cast on o0 ok")

	_ = toko0
	_ = entClient

	// if err := expectDisconnect(la, tok); err != nil {
	// 	log.Error().Err(err).Msg("disconnect")
	// 	return
	// }
	// log.Info().Msg("disconnect ok")
	log.Info().Msg("integration ok")
}
