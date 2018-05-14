package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudflare/golz4"
	"github.com/oklog/ulid"
	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01/dto"
	"github.com/elojah/services"
)

func send(cfg Config) {

	id, _ := ulid.Parse("01CDFNRKJJK8PBQV3DGHRM99BM")

	ack := [16]byte(id)

	conn, err := net.Dial("udp", "127.0.0.1:3400")
	if err != nil {
		fmt.Println("failed to establish connection", err)
		return
	}

	for i := 0; i < 10; i++ {
		// time.Sleep(1 * time.Second)
		go func() {
			msg :=
				dto.Message{
					Token: id,
					Action: dto.Attack{
						Source: id,
						Target: id,
					},
					ACK: &ack,
					TS:  time.Now().UnixNano(),
				}
			raw, err := msg.Marshal(nil)
			if err != nil {
				fmt.Println("error marshaling:", err)
				return
			}
			out := make([]byte, lz4.CompressBound(raw))
			n, err := lz4.Compress(raw, out)
			if err != nil {
				fmt.Println("failed to compress lz4", err)
				return
			}
			if _, err := conn.Write(out[:n]); err != nil {
				fmt.Println("write error")
			}
		}()
	}
}

// run services.
func run(prog string, filename string) {

	logger := logrus.NewEntry(logrus.New())
	logger = logger.WithField("app", filepath.Base(prog))

	launchers := services.Launchers{}

	cfg := Config{}
	cfgl := cfg.NewLauncher(Namespaces{
		API: "api",
	}, "api")
	launchers = append(launchers, cfgl)

	if err := launchers.Up(filename); err != nil {
		logger.WithField("filename", filename).Fatal(err.Error())
		return
	}

	logger.Info("api up")

	logger.Info("start sending")
	send(cfg)
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
