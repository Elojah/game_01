package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	abilitysrg "github.com/elojah/game_01/pkg/ability/srg"
	abilitysvc "github.com/elojah/game_01/pkg/ability/svc"
	accountsrg "github.com/elojah/game_01/pkg/account/srg"
	accountsvc "github.com/elojah/game_01/pkg/account/svc"
	entitysrg "github.com/elojah/game_01/pkg/entity/srg"
	entitysvc "github.com/elojah/game_01/pkg/entity/svc"
	eventsrg "github.com/elojah/game_01/pkg/event/srg"
	infrasrg "github.com/elojah/game_01/pkg/infra/srg"
	infrasvc "github.com/elojah/game_01/pkg/infra/svc"
	sectorsrg "github.com/elojah/game_01/pkg/sector/srg"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("exe", prog).Logger()

	launchers := services.Launchers{}

	// redis
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

	// Stores and applicatives
	accountStore := accountsrg.NewStore(rd)
	abilityStore := abilitysrg.NewStore(rd)
	entityStore := entitysrg.NewStore(rd)
	entityLRUStore := entitysrg.NewStore(rdlru)
	eventStore := eventsrg.NewStore(rd)
	infraStore := infrasrg.NewStore(rd)
	sectorStore := sectorsrg.NewStore(rd)

	// handler (https server)
	h := &handler{
		AccountStore:      accountStore,
		AccountTokenStore: accountStore,

		EntityStore:           entityLRUStore,
		EntityInventoryStore:  entityStore,
		EntityPCStore:         entityStore,
		EntityPCLeftStore:     entityStore,
		EntityPermissionStore: entityStore,
		EntityTemplateStore:   entityStore,

		EventQStore: eventStore,

		InfraSyncStore: infraStore,

		SectorStore:         sectorStore,
		SectorEntitiesStore: sectorStore,
		SectorStarterStore:  sectorStore,

		AbilityService: &abilitysvc.Service{
			AbilityStore:         abilityStore,
			AbilityStarterStore:  abilityStore,
			AbilityTemplateStore: abilityStore,
		},
		AccountTokenService: &accountsvc.TokenService{
			AccountStore:          accountStore,
			AccountTokenStore:     accountStore,
			EntityStore:           entityLRUStore,
			EntityPCStore:         entityStore,
			EntityPermissionStore: entityStore,
			EntityService: &entitysvc.Service{
				AbilityStore:          abilityStore,
				EntityStore:           entityLRUStore,
				EntityPermissionStore: entityStore,
				SectorEntitiesStore:   sectorStore,
				SequencerService: &infrasvc.SequencerService{
					InfraQSequencer: infraStore,
					InfraSequencer:  infraStore,
					InfraCore:       infraStore,
				},
			},
		},
		EntityPCService: &entitysvc.PCService{
			AbilityStore:      abilityStore,
			EntityPCStore:     entityStore,
			EntityPCLeftStore: entityStore,
		},
		InfraSequencerService: &infrasvc.SequencerService{
			InfraQSequencer: infraStore,
			InfraSequencer:  infraStore,
			InfraCore:       infraStore,
		},
		InfraRecurrerService: &infrasvc.RecurrerService{
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
				continue
			}
		case syscall.SIGKILL:
			if err := launchers.Down(); err != nil {
				log.Error().Err(err).Msg("failed to stop services")
				continue
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
