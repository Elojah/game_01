package main

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func logCmd(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)

	cmdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Str("cmd", args[0]).Msg("failed to pipe out")
		return err
	}

	go func() {
		r := bufio.NewReader(cmdout)
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				log.Error().Err(err).Msg("failed to read")
				return
			}
			log.Info().Msg(s)
		}
	}()

	return nil
}

func main() {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", os.Args[0]).Logger()

	cmds := [][]string{
		[]string{"game_sync", "-c", "configs/config_sync.json"},
		[]string{"game_core", "-c", "configs/config_core.json"},
		[]string{"game_api", "-c", "configs/config_api.json"},
		[]string{"game_auth", "-c", "configs/config_auth.json"},
		[]string{"game_revoker", "-c", "configs/config_revoker.json"},
		[]string{"game_tool", "-c", "configs/config_tool.json"},
	}

	for _, args := range cmds {
		if err := logCmd(args...); err != nil {
			log.Error().Err(err).Msg("failed to start")
		}
		log.Info().Msg("integration up")
	}

	log.Info().Msg("integration up")

	select {}
}
