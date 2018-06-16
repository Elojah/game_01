package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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
	launchers = append(launchers, rdl)
	rdx := redisx.NewService(&rd)

	// handler (https server)
	h := handler{}
	hl := h.NewLauncher(Namespaces{
		Tool: "tool",
	}, "tool")
	launchers = append(launchers, hl)

	h.AbilityTemplateMapper = rdx
	h.AccountMapper = rdx
	h.EntityMapper = rdx
	h.EntityTemplateMapper = rdx
	h.SectorMapper = rdx

	http.HandleFunc("/ability", h.ability)
	http.HandleFunc("/entity", h.entity)
	http.HandleFunc("/sector", h.sector)

	http.HandleFunc("/ability/template", h.abilityTemplate)
	http.HandleFunc("/entity/template", h.entityTemplate)

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	log.Info().Msg("tool up")
	select {}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}
	run(args[0], args[1])
}
