package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/usecase/listener"
	redisx "github.com/elojah/game_01/storage/redis"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""

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

	h.AccountMapper = rdx
	h.CoreMapper = rdx
	h.EntitiesMapper = rdlrux
	h.EntityMapper = rdlrux
	h.L = listener.L{
		ListenerMapper:  rdx,
		QListenerMapper: rdx,
		CoreMapper:      rdx,
	}
	h.PCLeftMapper = rdx
	h.PCMapper = rdx
	h.PermissionMapper = rdx
	h.QListenerMapper = rdx
	h.QMapper = rdx
	h.QRecurrerMapper = rdx
	h.SectorMapper = rdx
	h.StarterMapper = rdx
	h.SyncMapper = rdx
	h.TemplateMapper = rdx
	h.TokenMapper = rdx

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
