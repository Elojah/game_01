package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	accountapp "github.com/elojah/game_01/pkg/account/app"
	accountsrg "github.com/elojah/game_01/pkg/account/srg"
	entityapp "github.com/elojah/game_01/pkg/entity/app"
	entitysrg "github.com/elojah/game_01/pkg/entity/srg"
	eventapp "github.com/elojah/game_01/pkg/event/app"
	eventsrg "github.com/elojah/game_01/pkg/event/srg"
	"github.com/elojah/mux"
	"github.com/elojah/mux/client"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("exe", prog).Logger()

	launchers := services.Launchers{}

	rd := &redis.Service{}
	rdl := rd.NewLauncher(redis.Namespaces{
		Redis: "redis",
	}, "redis")
	launchers.Add(rdl)

	// redis-lru
	rdlru := &redis.Service{}
	rdlrul := rdlru.NewLauncher(redis.Namespaces{
		Redis: "redis-lru",
	}, "redis-lru")
	launchers.Add(rdlrul)

	m := &mux.M{}
	muxl := m.NewLauncher(mux.Namespaces{
		M: "server",
	}, "server")
	launchers.Add(muxl)

	c := &client.C{}
	cl := c.NewLauncher(client.Namespaces{
		Client: "client",
	}, "client")
	launchers.Add(cl)

	// Stores and applicatives
	eventStore := eventsrg.NewStore(rdlru)
	eventApp := &eventapp.A{
		QStore:       eventStore,
		Store:        eventStore,
		TriggerStore: eventStore,
	}

	entityStore := entitysrg.NewStore(rd)
	entityLRUStore := entitysrg.NewStore(rdlru)
	accountStore := accountsrg.NewStore(rd)
	accountApp := &accountapp.A{
		Store:        accountStore,
		TokenStore:   accountStore,
		TokenHCStore: accountStore,
		Entity: &entityapp.A{
			Store:           entityLRUStore,
			PCStore:         entityStore,
			PermissionStore: entityStore,
		},
	}

	h := &handler{
		M:       m,
		C:       c,
		event:   eventApp,
		account: accountApp,
	}

	hl := h.NewLauncher(Namespaces{
		API: "api",
	}, "api")
	launchers.Add(hl)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	log.Info().Msg("api up")

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
			if err := launchers.Down(); err != nil {
				log.Error().Err(err).Msg("failed to stop services")
			}
		case syscall.SIGKILL:
			if err := launchers.Down(); err != nil {
				log.Error().Err(err).Msg("failed to stop services")
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
