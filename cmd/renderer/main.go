package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/faiface/pixel/pixelgl"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""

	launchers := services.Launchers{}

	a := app{}
	al := a.NewLauncher(Namespaces{
		App: "renderer",
	}, "renderer")
	launchers.Add(al)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	entityC := make(chan entity.E, 0)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var e entity.E
			if err := json.Unmarshal(scanner.Bytes(), &e); err != nil {
				log.Error().Err(err).Msg("failed to decode entity")
				continue
			}
			log.Info().Str("entity", e.ID.String()).Msg("received entity")
			entityC <- e
		}
	}()

	log.Info().Msg("renderer up")
	pixelgl.Run(func() { a.Start(entityC) })

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, syscall.SIGHUP)
	for sig := range cs {
		switch sig {
		case syscall.SIGHUP:
			launchers.Down()
			launchers.Up(filename)
		case syscall.SIGINT:
			launchers.Down()
			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}
	run(args[0], args[1])
}
