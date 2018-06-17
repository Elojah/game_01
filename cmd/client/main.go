package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/elojah/mux"
	"github.com/elojah/services"
)

// run services.
func run(prog string, filename string) {

	logger := logrus.NewEntry(logrus.New())
	logger = logger.WithField("app", filepath.Base(prog))

	launchers := services.Launchers{}

	cfg := Config{}
	cfgl := cfg.NewLauncher(Namespaces{
		App: "app",
	}, "app")
	launchers = append(launchers, cfgl)

	m := mux.M{}
	muxl := m.NewLauncher(mux.Namespaces{
		M: "server",
	}, "server")
	launchers = append(launchers, muxl)

	if err := launchers.Up(filename); err != nil {
		logger.WithField("filename", filename).Fatal(err.Error())
		return
	}

	logger.Info("client up")
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
