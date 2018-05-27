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
	tile38x "github.com/elojah/game_01/storage/tile38"
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

	t38 := redis.Service{}
	t38l := t38.NewLauncher(redis.Namespaces{
		Redis: "tile38",
	}, "tile38")
	launchers = append(launchers, t38l)
	t38x := tile38x.NewService(&t38)

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

	a.AbilityMapper = rdx
	a.AbilityFeedbackMapper = rdx
	a.AbilityTemplateMapper = rdx
	a.EntityMapper = rdx
	a.EntityPositionMapper = t38x
	a.EntityTemplateMapper = rdx
	a.EventMapper = rdx
	a.PCMapper = rdx
	a.PCLeftMapper = rdx
	a.PermissionMapper = rdx
	a.QEventMapper = nax
	a.QListenerMapper = nax
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
