package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01"
	scyllax "github.com/elojah/game_01/storage/scylla"
	"github.com/elojah/scylla"
	"github.com/elojah/services"
	"github.com/elojah/udp"
)

func route(mux *udp.Mux, cfg Config) {
	for _, r := range cfg.Resources {
		if r == "actor" {
			a := actor{}
			mux.Add("ACCR", a.CreateActor)
			mux.Add("ACUP", a.UpdateActor)
			mux.Add("ACDE", a.DeleteActor)
			mux.Add("ACLI", a.ListActor)
		}
	}
	mux.Dispatcher = dispatch
}

func dispatch(packet []byte) (string, error) {
	msg := game.Message{}
	if _, err := msg.Unmarshal(packet); err != nil {
		return "", err
	}
	switch msg.Val.(type) {
	case game.Create:
		return "create", nil
	case game.Update:
		return "update", nil
	}
	return "", errors.New("unknown type")
}

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

	route(&mux, cfg)
	logger.Info("api up")
	go func() { mux.Listen() }()
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
