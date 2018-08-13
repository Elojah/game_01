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
	entityapp "github.com/elojah/game_01/pkg/entity/app"
	entitystore "github.com/elojah/game_01/pkg/entity/storage"
	eventstore "github.com/elojah/game_01/pkg/event/storage"
	infraapp "github.com/elojah/game_01/pkg/infra/app"
	infrastore "github.com/elojah/game_01/pkg/infra/storage"
	sectorstore "github.com/elojah/game_01/pkg/sector/storage"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("exe", prog).Logger()

	launchers := services.Launchers{}

	// redis
	rd := redis.Service{}
	rdl := rd.NewLauncher(redis.Namespaces{
		Redis: "redis",
	}, "redis")
	launchers.Add(rdl)

	// redis-lru
	rdlru := redis.Service{}
	rdlrul := rdlru.NewLauncher(redis.Namespaces{
		Redis: "redis-lru",
	}, "redis-lru")
	launchers.Add(rdlrul)

	// Stores and applicatives
	accountStore := accountstore.NewService(rd)
	accountApp := accountapp.NewService(rd)
	entityStore := entitystore.NewService(rd)
	entityLRUStore := entitystore.NewService(rdlru)
	eventStore := eventstore.NewService(rd)
	infraApp := infraapp.NewService(rd)
	infraStore := infrastore.NewService(rd)
	sectorStore := sectorstore.NewService(rd)

	// handler (https server)
	h := handler{
		AccountStore:    accountStore,
		PCStore:         entityStore,
		PCLeftStore:     entityStore,
		PermissionStore: entityStore,
		TemplateStore:   entityStore,
		QStore:          eventStore,
		SyncStore:       infraStore,
		EntitiesStore:   sectorStore,
		StarterStore:    sectorStore,
		SectorStore:     sectorStore,
		ListenerService: infraapp.ListenerService{
			InfraQListenerStore: infraStore,
			InfraListenerStore:  infraStore,
			InfraCoreStore:      infraStore,
		},
		TokenService: accountapp.TokenService{
			AccountStore:          accountStore,
			AccountTokenStore:     accountStore,
			EntityStore:           entityLRUStore,
			EntityPCStore:         entityStore,
			EntityPermissionStore: entityStore,

			EntityService: entityapp.Service{
				EntityStore:           entityLRUStore,
				EntityPermissionStore: entityStore,
				SectorEntitiesStore:   entityStore,

				ListenerService: infraapp.ListenerService{
					InfraQListenerStore: infraStore,
					InfraListenerStore:  infraStore,
					InfraCoreStore:      infraStore,
				},
			},
			InfraRecurrerService: infraapp.RecurrerService{
				InfraQRecurrerStore: infraStore,
				InfraRecurrerStore:  infraStore,
				InfraSyncStore:      infraStore,
			},
		},
	}
	hl := h.NewLauncher(Namespaces{
		Auth: "auth",
	}, "auth")
	launchers.Add(hl)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	log.Info().Msg("auth up")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	for sig := range c {
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
