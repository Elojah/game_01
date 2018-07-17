package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/elojah/game_01/cmd/client/login"
	"github.com/elojah/game_01/cmd/client/window"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""
	log.Logger = log.With().Str("exe", prog).Logger()

	launchers := services.Launchers{}

	w := window.NewWindow()
	wl := w.NewLauncher(window.Namespaces{
		Window: "window",
	}, "window")
	launchers.Add(wl)

	lr := login.NewRenderer()
	lrl := lr.NewLauncher(login.Namespaces{
		Login: "login",
	}, "login")
	launchers.Add(lrl)

	sdl.Main(func() {
		if err := launchers.Up(filename); err != nil {
			log.Error().Err(err).Str("filename", filename).Msg("failed to start")
			return
		}
		log.Info().Msg("client up")
		// sdl.Do(func() { r.UnstackEvent() })
		launchers.Down()
	})
}

/*
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
*/

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}

	run(args[0], args[1])
}
