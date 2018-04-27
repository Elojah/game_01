package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudflare/golz4"
	"github.com/gocql/gocql"
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
	id := gocql.TimeUUID()
	id, _ = gocql.ParseUUID("d91cb620-47cf-11e8-bef2-000000000001")
	msg :=
		dto.Message{
			Token:  id,
			Action: nil,
			ACK:    nil,
			TS:     time.Now().Unix(),
		}

	raw, err := msg.Marshal(nil)
	if err != nil {
		fmt.Println("error marshaling:", err)
		return
	}

	conn, err := net.Dial("tcp", "127.0.0.1:3400")
	if err != nil {
		fmt.Println("failed to establish connection", err)
		return
	}

	out := make([]byte, lz4.CompressBound(raw))

	n, err := lz4.Compress(raw, out)
	if err != nil {
		fmt.Println("failed to compress lz4", err)
		return
	}

	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("send shit")
		conn.Write(out[:n])
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
