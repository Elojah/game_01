package main

import (
	"errors"

	"github.com/elojah/game_01"
	"github.com/elojah/udp"
)

func route(mux *udp.Mux, cfg Config) {
	for _, r := range cfg.Resources {
		if r == "actor" {
			a := actor{}
			mux.Add("actor_create", a.CreateActor)
			mux.Add("actor_update", a.UpdateActor)
			mux.Add("actor_delete", a.DeleteActor)
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
	case game.ActorCreate:
		return "actor_create", nil
	case game.ActorUpdate:
		return "actor_update", nil
	case game.ActorDelete:
		return "actor_delete", nil
	}
	return "", errors.New("unknown type")
}
