package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	abilitysrg "github.com/elojah/game_01/pkg/ability/srg"
	accountsrg "github.com/elojah/game_01/pkg/account/srg"
	entitysrg "github.com/elojah/game_01/pkg/entity/srg"
	entitysvc "github.com/elojah/game_01/pkg/entity/svc"
	eventsrg "github.com/elojah/game_01/pkg/event/srg"
	eventsvc "github.com/elojah/game_01/pkg/event/svc"
	infrasrg "github.com/elojah/game_01/pkg/infra/srg"
	itemsrg "github.com/elojah/game_01/pkg/item/srg"
	sectorsrg "github.com/elojah/game_01/pkg/sector/srg"
	sectorsvc "github.com/elojah/game_01/pkg/sector/svc"
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
	eventStore := eventsrg.NewStore(rdlru)
	abilityStore := abilitysrg.NewStore(rd)
	abilityLRUStore := abilitysrg.NewStore(rdlru)
	accountStore := accountsrg.NewStore(rd)
	entityStore := entitysrg.NewStore(rd)
	entityLRUStore := entitysrg.NewStore(rdlru)
	infraStore := infrasrg.NewStore(rd)
	itemStore := itemsrg.NewStore(rd)
	sectorStore := sectorsrg.NewStore(rd)

	// main service
	svc := &service{
		AbilityStore:         abilityStore,
		AbilityTemplateStore: abilityStore,
		AbilityFeedbackStore: abilityLRUStore,

		TokenStore: accountStore,

		EntityStore:           entityLRUStore,
		EntityTemplateStore:   entityStore,
		EntityPermissionStore: entityStore,
		EntityInventoryService: &entitysvc.InventoryService{
			EntityInventoryStore:   entityLRUStore,
			EntityMRInventoryStore: entityStore,
		},
		EntityPermissionService: entitysvc.PermissionService{
			EntityPermissionStore: entityStore,
		},
		EntitySpawnStore: entityStore,

		QSequencerStore: infraStore,
		CoreStore:       infraStore,

		EventQStore: eventStore,
		EventStore:  eventStore,
		EventTriggerService: &eventsvc.TriggerService{
			TriggerStore: eventStore,
			Store:        eventStore,
			QStore:       eventStore,
		},

		EntitiesStore: sectorStore,
		ItemStore:     itemStore,
		ItemLootStore: itemStore,
		SectorStore:   sectorStore,
		SectorService: &sectorsvc.Service{
			SectorEntitiesStore: sectorStore,
			SectorStore:         sectorStore,
		},
	}
	svcl := svc.NewLauncher(Namespaces{
		Service: "service",
	}, "service")
	launchers.Add(svcl)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	log.Info().Msg("core up")
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
