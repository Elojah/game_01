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

	na.Flush()

	a.AbilityMapper = rdx
	a.FeedbackMapper = rdx
	a.AbilityTemplateMapper = rdx
	a.EntityMapper = rdlrux
	a.EntityTemplateMapper = rdx
	a.EventMapper = rdx
	a.PermissionMapper = rdx
	a.QMapper = nax
	a.QListenerMapper = nax
	a.SectorMapper = rdx
	a.EntitiesMapper = rdlrux
	a.SubscriptionMapper = nax
	a.TokenMapper = rdx

	go func() { a.Start() }()
	log.Info().Msg("core up")

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
