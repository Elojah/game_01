package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	accountapp "github.com/elojah/game_01/pkg/account/app"
	accountstore "github.com/elojah/game_01/pkg/account/storage"
	entitystore "github.com/elojah/game_01/pkg/entity/storage"
	eventstore "github.com/elojah/game_01/pkg/event/storage"
	infraapp "github.com/elojah/game_01/pkg/infra/app"
	infrastore "github.com/elojah/game_01/pkg/infra/storage"
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
	eventStore := eventstore.NewService(rd)
	accountStore := accountstore.NewService(rd)
	entityStore := entitystore.NewService(rd)
	entityLRUStore := entitystore.NewService(rdlru)
	infraStore := infrastore.NewService(rd)

	h := &handler{
		M:      m,
		C:      c,
		QStore: eventStore,
		TokenService: accountapp.TokenService{
			AccountStore:      accountStore,
			AccountTokenStore: accountStore,
			EntityStore:       entityLRUStore,
			EntityPCStore:     entityStore,
			InfraRecurrerService: infraapp.RecurrerService{
				InfraQRecurrerStore: infraStore,
				InfraRecurrerStore:  infraStore,
				InfraSyncStore:      infraStore,
			},
			EntityPermissionStore: entityStore,
		},
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
