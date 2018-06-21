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

	// redis
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

	// handler (https server)
	h := handler{}
	hl := h.NewLauncher(Namespaces{
		Auth: "auth",
	}, "auth")
	launchers = append(launchers, hl)

	h.AccountMapper = rdx
	h.EntityMapper = rdx
	h.EntityTemplateMapper = rdx
	h.PCMapper = rdx
	h.PCLeftMapper = rdx
	h.QEventMapper = nax
	h.QListenerMapper = nax
	h.QRecurrerMapper = nax
	h.SectorEntitiesMapper = rdx
	h.TokenMapper = rdx

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	log.Info().Msg("auth up")
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
