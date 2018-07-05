package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/usecase/listener"
	"github.com/elojah/game_01/pkg/usecase/recurrer"
	"github.com/elojah/game_01/pkg/usecase/token"
	redisx "github.com/elojah/game_01/storage/redis"
	"github.com/elojah/mux"
	"github.com/elojah/mux/client"
	"github.com/elojah/redis"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""

	launchers := services.Launchers{}

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

	m := mux.M{}
	muxl := m.NewLauncher(mux.Namespaces{
		M: "server",
	}, "server")
	launchers.Add(muxl)

	c := client.C{}
	cl := c.NewLauncher(client.Namespaces{
		Client: "client",
	}, "client")
	launchers.Add(cl)

	h := handler{
		M:       &m,
		C:       &c,
		QMapper: rdx,
		T: token.T{
			AccountMapper: rdx,
			TokenMapper:   rdx,
			EntityMapper:  rdlrux,
			PCMapper:      rdx,
			L: listener.L{
				QListenerMapper: rdx,
				ListenerMapper:  rdx,
				CoreMapper:      rdx,
			},
			R: recurrer.R{
				QRecurrerMapper: rdx,
				RecurrerMapper:  rdx,
				SyncMapper:      rdx,
			},
			PermissionMapper: rdx,
			EntitiesMapper:   rdlrux,
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
