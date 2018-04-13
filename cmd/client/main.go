package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/oklog/ulid"
	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01/dto"
	"github.com/elojah/services"
	"github.com/elojah/udp"
)

func route(mux *udp.Mux, cfg Config) {
	// for _,r_ := range cfg.Resources {
	// 	// Add gamestate handling here
	// }
}

func send(mux *udp.Mux, cfg Config) {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader)
	packetID := ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader)
	msg :=
		dto.Message{
			Token: id,
			Position: &dto.Vec3{
				X: 190.098,
				Y: 34.9,
				Z: 98.27,
			},
			Action: nil,
			ACK:    nil,
		}

	raw, err := msg.Marshal(nil)
	if err != nil {
		fmt.Println("error marshaling:", err)
		return
	}

	for n := 0; n < 60; n++ {
		go func() { mux.Send(udp.Packet{ID: packetID, Source: nil, Data: raw}, "127.0.0.1:3400") }()
	}
}

// run services.
func run(prog string, filename string) {

	logger := logrus.NewEntry(logrus.New())
	logger = logger.WithField("app", filepath.Base(prog))

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

	route(&mux, cfg)
	logger.Info("api up")

	logger.Info("start sending")
	send(&mux, cfg)
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
