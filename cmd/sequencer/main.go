package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	natsx "github.com/elojah/game_01/storage/nats"
	redisx "github.com/elojah/game_01/storage/redis"
	"github.com/elojah/nats"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""

	launchers := services.Launchers{}

	rd := redis.Service{}
	rdl := rd.NewLauncher(redis.Namespaces{
		Redis: "redis",
	}, "redis")
	launchers = append(launchers, rdl)
	rdx := redisx.NewService(&rd)

	na := nats.Service{}
	nal := na.NewLauncher(nats.Namespaces{
		Nats: "nats",
	}, "nats")
	launchers = append(launchers, nal)
	nax := natsx.NewService(&na)

	a := app{}
	al := a.NewLauncher(Namespaces{
		App: "app",
	}, "app")
	launchers = append(launchers, al)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	na.Flush()

	a.EntityService = rdx
	a.EventService = rdx
	a.QEventService = nax
	a.SubscriptionService = nax

	go func() { a.Start() }()
	log.Info().Msg("sequencer up")

	c := make(chan os.Signal, 0)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case _ = <-c:
			a.Close()
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
