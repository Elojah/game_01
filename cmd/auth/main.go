package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	redisx "github.com/elojah/game_01/pkg/storage/redis"
	"github.com/elojah/game_01/pkg/usecase/listener"
	"github.com/elojah/game_01/pkg/usecase/recurrer"
	"github.com/elojah/game_01/pkg/usecase/token"
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
	rdx := redisx.NewService(&rd)

	// redis-lru
	rdlru := redis.Service{}
	rdlrul := rdlru.NewLauncher(redis.Namespaces{
		Redis: "redis-lru",
	}, "redis-lru")
	launchers.Add(rdlrul)
	rdlrux := redisx.NewService(&rdlru)

	// handler (https server)
	h := handler{}
	hl := h.NewLauncher(Namespaces{
		Auth: "auth",
	}, "auth")
	launchers.Add(hl)

	l := listener.L{
		QListenerService: rdx,
		ListenerService:  rdx,
		CoreService:      rdx,
	}
	h.AccountService = rdx
	h.EntitiesService = rdlrux
	h.PCService = rdx
	h.PCLeftService = rdx
	h.PermissionService = rdx
	h.L = l
	h.QService = rdx
	h.SectorService = rdx
	h.StarterService = rdx
	h.T = token.T{
		L: l,
		R: recurrer.R{
			QRecurrerService: rdx,
			RecurrerService:  rdx,
			SyncService:      rdx,
		},
		AccountService:    rdx,
		EntityService:     rdlrux,
		TokenService:      rdx,
		PCService:         rdx,
		PermissionService: rdx,
		EntitiesService:   rdlrux,
	}
	h.TemplateService = rdx

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
