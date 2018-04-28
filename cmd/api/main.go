package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	influxx "github.com/elojah/game_01/storage/influx"
	scyllax "github.com/elojah/game_01/storage/scylla"
	"github.com/elojah/influx"
	"github.com/elojah/mux"
	"github.com/elojah/scylla"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	zerolog.TimeFieldFormat = ""

	launchers := services.Launchers{}

	sc := scylla.Service{}
	scl := sc.NewLauncher(scylla.Namespaces{
		Scylla: "scylla",
	}, "scylla")
	launchers = append(launchers, scl)
	scx := scyllax.NewService(&sc)

	in := influx.Service{}
	inl := in.NewLauncher(influx.Namespaces{
		Influx: "influx",
	}, "influx")
	launchers = append(launchers, inl)
	inx := influxx.NewService(&in)

	m := mux.M{}
	muxl := m.NewLauncher(mux.Namespaces{
		M: "server",
	}, "server")
	launchers = append(launchers, muxl)

	cfg := Config{}
	cfgl := cfg.NewLauncher(Namespaces{
		API: "api",
	}, "api")
	launchers = append(launchers, cfgl)

	_ = scx

	if err := launchers.Up(filename); err != nil {
		log.Error().Err(err).Str("filename", filename).Msg("failed to start")
		return
	}

	h := handler{}
	h.Services = game.NewServices()
	h.Config = cfg
	h.TokenService = scx
	h.ActorService = inx
	h.Route(&m, cfg)

	go func() { m.Listen() }()
	log.Info().Msg("api up")
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
