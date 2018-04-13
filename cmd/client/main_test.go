package main

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/elojah/services"
	"github.com/elojah/udp"
)

func BenchmarkSend(b *testing.B) {
	prog := "test"
	filename := "../../bin/config_client.json"

	logger := logrus.NewEntry(logrus.New())
	logger = logger.WithField("app", prog)

	launchers := services.Launchers{}

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

	if err := launchers.Up(filename); err != nil {
		logger.WithField("filename", filename).Fatal(err.Error())
		return
	}
	defer muxl.Down(services.Configs{})

	route(&mux, cfg)
	logger.Info("api up")

	logger.Info("start sending")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			send(&mux, cfg)
		}
	})
}
