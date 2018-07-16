package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", prog).Logger()

	launchers := services.Launchers{}

	// main app
	a := app{}
	al := a.NewLauncher(Namespaces{
		App: "booter",
	}, "booter")
	launchers.Add(al)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	log.Info().Msg("booter up")
	go func() {
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
	}()
	a.Start()
	launchers.Down()
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}
	run(args[0], args[1])
}
