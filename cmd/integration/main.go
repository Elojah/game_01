package main

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", os.Args[0]).Logger()

	sync := exec.Command("game_sync", "-c", "configs/config_sync.json")
	core := exec.Command("game_core", "-c", "configs/config_core.json")
	api := exec.Command("game_api", "-c", "configs/config_api.json")
	auth := exec.Command("game_auth", "-c", "configs/config_auth.json")
	tool := exec.Command("game_tool", "-c", "configs/config_tool.json")

	syncout, err := sync.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Msg("failed to pipe out sync")
		return
	}
	coreout, err := core.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Msg("failed to pipe out core")
		return
	}
	apiout, err := api.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Msg("failed to pipe out api")
		return
	}
	authout, err := auth.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Msg("failed to pipe out auth")
		return
	}
	toolout, err := tool.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Msg("failed to pipe out tool")
		return
	}

	go func() {
		r := bufio.NewReader(syncout)
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				log.Error().Err(err).Msg("failed to read sync")
				return
			}
			log.Info().Msg(s)
		}
	}()
	go func() {
		r := bufio.NewReader(coreout)
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				log.Error().Err(err).Msg("failed to read core")
				return
			}
			log.Info().Msg(s)
		}
	}()
	go func() {
		r := bufio.NewReader(apiout)
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				log.Error().Err(err).Msg("failed to read api")
				return
			}
			log.Info().Msg(s)
		}
	}()
	go func() {
		r := bufio.NewReader(authout)
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				log.Error().Err(err).Msg("failed to read auth")
				return
			}
			log.Info().Msg(s)
		}
	}()
	go func() {
		r := bufio.NewReader(toolout)
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				log.Error().Err(err).Msg("failed to read tool")
				return
			}
			log.Info().Msg(s)
		}
	}()

}
