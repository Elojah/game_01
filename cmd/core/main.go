package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	abilityapp "github.com/elojah/game_01/pkg/ability/app"
	abilitysrg "github.com/elojah/game_01/pkg/ability/srg"
	accountapp "github.com/elojah/game_01/pkg/account/app"
	accountsrg "github.com/elojah/game_01/pkg/account/srg"
	entityapp "github.com/elojah/game_01/pkg/entity/app"
	entitysrg "github.com/elojah/game_01/pkg/entity/srg"
	eventapp "github.com/elojah/game_01/pkg/event/app"
	eventsrg "github.com/elojah/game_01/pkg/event/srg"
	infraapp "github.com/elojah/game_01/pkg/infra/app"
	infrasrg "github.com/elojah/game_01/pkg/infra/srg"
	itemapp "github.com/elojah/game_01/pkg/item/app"
	itemsrg "github.com/elojah/game_01/pkg/item/srg"
	sectorapp "github.com/elojah/game_01/pkg/sector/app"
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

	abilityStore := abilitysrg.NewStore(rd)
	abilityLRUStore := abilitysrg.NewStore(rdlru)
	abilityApp := &abilityapp.A{
		FeedbackStore: abilityLRUStore,
		StarterStore:  abilityStore,
		Store:         abilityStore,
		TemplateStore: abilityStore,
	}

	infraStore := infrasrg.NewStore(rd)
	sequencerApp := &infraapp.SequencerApp{
		QSequencerStore: infraStore,
		SequencerStore:  infraStore,
		CoreStore:       infraStore,
	}

	entityStore := entitysrg.NewStore(rd)
	entityLRUStore := entitysrg.NewStore(rdlru)
	sectorStore := sectorsrg.NewStore(rd)
	entityApp := &entityapp.A{
		InventoryStore:      entityLRUStore,
		MRInventoryStore:    entityStore,
		PCLeftStore:         entityStore,
		PCStore:             entityStore,
		PermissionStore:     entityStore,
		SpawnStore:          entityStore,
		Store:               entityLRUStore,
		TemplateStore:       entityStore,
		AbilityStore:        abilityStore,
		SectorEntitiesStore: sectorStore,
		Sequencer:           sequencerApp,
	}

	accountStore := accountsrg.NewStore(rd)
	accountApp := &accountapp.A{
		Store:        accountStore,
		TokenStore:   accountStore,
		TokenHCStore: accountStore,
		Entity:       entityApp,
	}

	eventStore := eventsrg.NewStore(rdlru)
	eventApp := &eventapp.A{
		QStore:       eventStore,
		Store:        eventStore,
		TriggerStore: eventStore,
	}

	itemStore := itemsrg.NewStore(rd)
	itemApp := &itemapp.A{
		Store: itemStore,
	}

	sectorApp := &sectorapp.A{
		Store:         sectorStore,
		EntitiesStore: sectorStore,
	}

	// main service
	svc := &service{
		ability:   abilityApp,
		account:   accountApp,
		entity:    entityApp,
		event:     eventApp,
		item:      itemApp,
		sector:    sectorApp,
		sequencer: sequencerApp,
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
