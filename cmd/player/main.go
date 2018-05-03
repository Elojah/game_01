package main

import (
	"fmt"
	"os"

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

	cfg := Config{}
	cfgl := cfg.NewLauncher(Namespaces{
		Player: "player",
	}, "player")
	launchers = append(launchers, cfgl)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	a := app{}
	a.Config = cfg
	a.EventService = nax
	a.EntityService = rdx

	go func() { a.Start() }()
	log.Info().Msg("player up")
	select {}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}
	run(args[0], args[1])
}
