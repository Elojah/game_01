package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/cmd/client/handler"
	"github.com/elojah/game_01/cmd/client/reader"
	"github.com/elojah/mux"
	"github.com/elojah/mux/client"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""

	launchers := services.Launchers{}

	m := mux.M{}
	muxl := m.NewLauncher(mux.Namespaces{
		M: "entity",
	}, "entity")
	launchers.Add(muxl)

	ma := mux.M{}
	mauxl := ma.NewLauncher(mux.Namespaces{
		M: "ack",
	}, "ack")
	launchers.Add(mauxl)

	c := client.C{}
	cl := c.NewLauncher(client.Namespaces{
		Client: "client",
	}, "client")
	launchers.Add(cl)

	h := handler.H{
		M: &m,
	}
	hl := h.NewLauncher(handler.Namespaces{
		Handler: "handler",
	}, "handler")
	launchers.Add(hl)
	h.M.Handler = h.HandleEntity

	ha := handler.H{
		M: &ma,
	}
	hal := ha.NewLauncher(handler.Namespaces{
		Handler: "handler",
	}, "handler")
	launchers.Add(hal)
	ha.M.Handler = ha.HandleACK

	rd := reader.NewReader(&c, ha.ACK)
	rdl := rd.NewLauncher(reader.Namespaces{
		Reader: "reader",
	}, "reader")
	launchers.Add(rdl)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	log.Info().Msg("client up")
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
