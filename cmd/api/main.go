package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	scyllax "github.com/elojah/game_01/storage/scylla"
	"github.com/elojah/scylla"
	"github.com/elojah/services"
	"github.com/elojah/udp"
)

// run services.
func run(prog string, filename string) {

	logger := logrus.NewEntry(logrus.New())
	logger = logger.WithField("app", filepath.Base(prog))

	launchers := services.Launchers{}

	sc := scylla.Service{}
	scl := sc.NewLauncher(scylla.Namespaces{
		Scylla: "scylla",
	}, "scylla")
	launchers = append(launchers, scl)
	scx := scyllax.NewService(&sc)

	mux := udp.Mux{}
	mux.Entry = logger
	muxl := mux.NewLauncher(udp.Namespaces{
		UDP: "server",
	}, "server")
	launchers = append(launchers, muxl)

	cfg := Config{}
	cfgl := cfg.NewLauncher(Namespaces{
		API: "api",
	}, "api")
	launchers = append(launchers, cfgl)

	_ = scx

	if err := launchers.Up(filename); err != nil {
		logger.WithField("filename", filename).Fatal(err.Error())
		return
	}

	h := handler{}
	h.Entry = logger
	h.Route(&mux, cfg)
	go func() { mux.Listen() }()
	logger.Info("api up")
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: ./%s configfile\n", args[0])
		return
	}
	run(args[0], args[1])
	select {}
}
