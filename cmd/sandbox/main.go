package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/cmd/sandbox/reader"
	"github.com/elojah/game_01/cmd/sandbox/ui"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
	"github.com/elojah/mux/client"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("exe", prog).Logger()

	launchers := services.Launchers{}

	// Handle entity DTO from sync
	m := mux.M{}
	muxl := m.NewLauncher(mux.Namespaces{
		M: "entity",
	}, "entity")
	launchers.Add(muxl)
	h := handler{
		M: &m,
	}
	hl := h.NewLauncher(NamespacesHandler{
		Handler: "handler",
	}, "handler")
	launchers.Add(hl)
	h.M.Handler = h.handleDTO

	// Handle ACK from api
	ma := mux.M{}
	mauxl := ma.NewLauncher(mux.Namespaces{
		M: "ack",
	}, "ack")
	launchers.Add(mauxl)

	ha := handler{
		M: &ma,
	}
	hal := ha.NewLauncher(NamespacesHandler{
		Handler: "handler",
	}, "handler")
	launchers.Add(hal)
	ha.M.Handler = ha.handleACK

	c := client.C{}
	cl := c.NewLauncher(client.Namespaces{
		Client: "client",
	}, "client")
	launchers.Add(cl)

	ack := make(chan gulid.ID)
	rd := reader.New(&c, ha.ACK)
	rdl := rd.NewLauncher(reader.Namespaces{
		Reader: "reader",
	}, "reader")
	launchers.Add(rdl)

	/*
		UI
	*/

	t := ui.Term{}
	tl := t.NewLauncher(ui.Namespaces{
		UI: "ui",
	}, "ui")
	launchers.Add(tl)
	log.Info().Msg("sandbox up")

	/*
		Run services
	*/

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, syscall.SIGHUP)
	for sig := range cs {
		switch sig {
		case syscall.SIGHUP:
			if err := launchers.Down(); err != nil {
				log.Error().Err(err).Msg("failed to stop services")
				continue
			}
			if err := launchers.Up(filename); err != nil {
				log.Error().Err(err).Str("filename", filename).Msg("failed to start services")
			}
		case syscall.SIGINT:
			close(ack)
			if err := launchers.Down(); err != nil {
				log.Error().Err(err).Msg("failed to stop services")
				continue
			}
		case syscall.SIGKILL:
			close(ack)
			if err := launchers.Down(); err != nil {
				log.Error().Err(err).Msg("failed to stop services")
				continue
			}
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
