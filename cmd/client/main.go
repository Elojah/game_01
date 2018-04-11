package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01"
	"github.com/elojah/services"
	"github.com/elojah/udp"
)

func route(mux *udp.Mux, cfg Config) {
	// for _,r_ := range cfg.Resources {
	// 	// Add gamestate handling here
	// }
}

func send(mux *udp.Mux, cfg Config) {
	b := flatbuffers.NewBuilder(0)
	game.ActorStart(b)
	game.ActorAddHp(b, 100)
	game.ActorAddMp(b, 100)
	actor := game.ActorEnd(b)
	b.Finish(actor)

	// b.Reset()
	// b = flatbuffers.NewBuilder(0)
	game.ActorCreateStartActorsVector(b, 1)
	b.PrependUOffsetT(actor)
	actors := b.EndVector(1)

	game.ActorCreateStart(b)
	game.ActorCreateAddActors(b, actors)
	actorCreate := game.ActorCreateEnd(b)
	b.Finish(actorCreate)

	for {
		// mux.Write([]byte("A"), "127.0.0.1:3400")
		mux.Send("test", b.Bytes[b.Head():], "127.0.0.1:3400")
		time.Sleep(time.Millisecond * time.Duration(cfg.TickRate))
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
