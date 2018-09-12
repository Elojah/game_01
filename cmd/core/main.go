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
	eventsrg "github.com/elojah/game_01/pkg/event/srg"
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
	eventStore := eventsrg.NewStore(rd)
	abilityStore := abilitysrg.NewStore(rd)
	accountStore := accountsrg.NewStore(rd)
	entityStore := entitysrg.NewStore(rd)
	entityLRUStore := entitysrg.NewStore(rdlru)
	infraStore := infrasrg.NewStore(rd)
	sectorStore := sectorsrg.NewStore(rd)

	// main app
	a := &app{
		AbilityStore:         abilityStore,
		AbilityTemplateStore: abilityStore,
		FeedbackStore:        abilityStore,

		TokenStore: accountStore,

		EntityStore:         entityLRUStore,
		EntityTemplateStore: entityStore,
		PermissionStore:     entityStore,

		QSequencerStore: infraStore,
		CoreStore:       infraStore,

		EventQStore: eventStore,
		EventStore:  eventStore,

		EntitiesStore: sectorStore,
		SectorStore:   sectorStore,
	}
	al := a.NewLauncher(Namespaces{
		App: "app",
	}, "app")
	launchers.Add(al)

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
			launchers.Down()
			launchers.Up(filename)
		case syscall.SIGINT:
			launchers.Down()
		case syscall.SIGKILL:
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
