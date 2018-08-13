package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	accountstore "github.com/elojah/game_01/pkg/account/storage"
	accountsvc "github.com/elojah/game_01/pkg/account/svc"
	entitystore "github.com/elojah/game_01/pkg/entity/storage"
	entitysvc "github.com/elojah/game_01/pkg/entity/svc"
	eventstore "github.com/elojah/game_01/pkg/event/storage"
	infrastore "github.com/elojah/game_01/pkg/infra/storage"
	infrasvc "github.com/elojah/game_01/pkg/infra/svc"
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
	accountStore := accountstore.NewStore(rd)
	entityStore := entitystore.NewStore(rd)
	entityLRUStore := entitystore.NewStore(rdlru)
	eventStore := eventstore.NewStore(rd)
	infraStore := infrastore.NewStore(rd)
	sectorStore := sectorstore.NewStore(rd)

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
		TokenService: accountsvc.TokenService{
			Account:          accountStore,
			AccountToken:     accountStore,
			Entity:           entityLRUStore,
			EntityPC:         entityStore,
			EntityPermission: entityStore,
			EntityService: entitysvc.Service{
				Entity:           entityLRUStore,
				EntityPermission: entityStore,
				SectorEntities:   entityStore,
				ListenerService: infrasvc.ListenerService{
					InfraQListener: infraStore,
					InfraListener:  infraStore,
					InfraCore:      infraStore,
				},
			},
			InfraRecurrerService: infrasvc.RecurrerService{
				InfraQRecurrer: infraStore,
				InfraRecurrer:  infraStore,
				InfraSync:      infraStore,
			},
		},
		ListenerService: infrasvc.ListenerService{
			InfraQListener: infraStore,
			InfraListener:  infraStore,
			InfraCore:      infraStore,
		},
		RecurrerService: infrasvc.RecurrerService{
			InfraQRecurrer: infraStore,
			InfraRecurrer:  infraStore,
			InfraSync:      infraStore,
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
