package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	abilitysrg "github.com/elojah/game_01/pkg/ability/srg"
	accountapp "github.com/elojah/game_01/pkg/account/app"
	accountsrg "github.com/elojah/game_01/pkg/account/srg"
	entityapp "github.com/elojah/game_01/pkg/entity/app"
	entitysrg "github.com/elojah/game_01/pkg/entity/srg"
	infraapp "github.com/elojah/game_01/pkg/infra/app"
	infrasrg "github.com/elojah/game_01/pkg/infra/srg"
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

	infraStore := infrasrg.NewStore(rd)
	sequencerApp := &infraapp.SequencerApp{
		QSequencerStore: infraStore,
		SequencerStore:  infraStore,
		CoreStore:       infraStore,
	}

	entityStore := entitysrg.NewStore(rd)
	entityLRUStore := entitysrg.NewStore(rdlru)
	sectorStore := sectorsrg.NewStore(rd)
	abilityStore := abilitysrg.NewStore(rd)
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

	// main service
	svc := &service{
		account: accountApp,
	}
	svcl := svc.NewLauncher(Namespaces{
		Service: "service",
	}, "service")
	launchers.Add(svcl)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	log.Info().Msg("revoker up")
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
