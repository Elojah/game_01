package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	natsx "github.com/elojah/game_01/storage/nats"
	redisx "github.com/elojah/game_01/storage/redis"
	"github.com/elojah/mux"
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

	// redis-lru
	rdlru := redis.Service{}
	rdlrul := rdlru.NewLauncher(redis.Namespaces{
		Redis: "redis-lru",
	}, "redis-lru")
	launchers = append(launchers, rdlrul)
	rdlrux := redisx.NewService(&rdlru)

	// nats
	na := nats.Service{}
	nal := na.NewLauncher(nats.Namespaces{
		Nats: "nats",
	}, "nats")
	launchers = append(launchers, nal)
	nax := natsx.NewService(&na)

	// mux
	m := mux.M{}
	muxl := m.NewLauncher(mux.Namespaces{
		M: "server",
	}, "server")
	launchers = append(launchers, muxl)

	// main app
	a := app{}
	al := a.NewLauncher(Namespaces{
		App: "app",
	}, "app")
	launchers = append(launchers, al)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	a.M = &m
	a.EntityMapper = rdlrux
	a.QMapper = nax
	a.QRecurrerMapper = nax
	a.EntitiesMapper = rdlrux
	a.SectorMapper = rdx
	a.SubscriptionMapper = nax
	a.SyncMapper = rdx
	a.TokenMapper = rdx

	go a.Start()
	log.Info().Msg("sync up")
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
