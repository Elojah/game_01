package main

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/elojah/services"
)

func BenchmarkSend(b *testing.B) {
	prog := "test"
	filename := "../../bin/config_client.json"

	logger := logrus.NewEntry(logrus.New())
	logger = logger.WithField("app", prog)

	launchers := services.Launchers{}

	cfg := Config{}
	cfgl := cfg.NewLauncher(Namespaces{
		App: "app",
	}, "app")
	launchers = append(launchers, cfgl)

	if err := launchers.Up(filename); err != nil {
		logger.WithField("filename", filename).Fatal(err.Error())
		return
	}

	logger.Info("app up")

	logger.Info("start sending")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			send(cfg)
		}
	})
}
