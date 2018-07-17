package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/elojah/game_01/cmd/client/booter"
	"github.com/elojah/game_01/cmd/client/handler"
	"github.com/elojah/game_01/cmd/client/renderer"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/usecase/token"
	"github.com/elojah/mux"
	"github.com/elojah/mux/client"
	"github.com/elojah/services"
)

// run services.
func run(filename string, t token.T, e entity.E) {

	launchers := services.Launchers{}

	c := client.C{}
	cl := c.NewLauncher(client.Namespaces{
		Client: "client",
	}, "client")
	launchers.Add(cl)

	me := mux.M{}
	meuxl := me.NewLauncher(mux.Namespaces{
		M: "entity",
	}, "entity")
	launchers.Add(meuxl)

	he := handler.NewHandler(&me)
	hel := he.NewLauncher(handler.Namespaces{
		Handler: "handler",
	}, "handler")
	launchers.Add(hel)
	he.M.Handler = he.HandleEntity

	ma := mux.M{}
	mauxl := ma.NewLauncher(mux.Namespaces{
		M: "ack",
	}, "ack")
	launchers.Add(mauxl)

	ha := handler.NewHandler(&ma)
	hal := ha.NewLauncher(handler.Namespaces{
		Handler: "handler",
	}, "handler")
	launchers.Add(hal)
	ha.M.Handler = ha.HandleACK

	r := renderer.NewRenderer(&c, ha.ACK, he.Entity)
	rl := r.NewLauncher(renderer.Namespaces{
		Renderer: "renderer",
	}, "renderer")
	launchers.Add(rl)

	sdl.Main(func() {
		if err := launchers.Up(filename); err != nil {
			log.Error().Err(err).Str("filename", filename).Msg("failed to start")
			return
		}
		log.Info().Msg("client up")
		sdl.Do(func() { r.UnstackEvent() })
		launchers.Down()
	})
}

func boot(filename string) (token.T, entity.E, error) {
	launchers := services.Launchers{}

	b := booter.B{}
	bl := b.NewLauncher(booter.Namespaces{
		Booter: "booter",
	}, "booter")
	launchers.Add(bl)

	var t token.T
	var e entity.E
	sdl.Main(func() {
		if err := launchers.Up(filename); err != nil {
			log.Error().Err(err).Str("filename", filename).Msg("failed to start")
			return
		}
		log.Info().Msg("booter up")
		sdl.Do(func() { b.UnstackEvent() })
		launchers.Down()
	})

	return t, e, errors.New("WIP TMP")
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}
	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", args[0]).Logger()

	t, e, err := boot(args[1])
	if err != nil {
		fmt.Println("Failed to boot game:", err.Error())
		return
	}
	run(args[1], t, e)
}
