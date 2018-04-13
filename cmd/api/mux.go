package main

import (
	"github.com/elojah/game_01/dto"
	"github.com/elojah/udp"
)

func route(mux *udp.Mux, cfg Config) {
	mux.Handler = dispatch
}

func dispatch(packet udp.Packet) error {
	msg := dto.Message{}
	if _, err := msg.Unmarshal(packet.Data); err != nil {
		return err
	}
	return nil
}
