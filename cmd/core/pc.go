package main

import (
	"math/rand"

	"github.com/elojah/game_01"
)

const (
	characterKey = "pc"
)

func (a *app) CreatePC(event game.Event) error {

	spc := event.Action.(game.SetPC)
	_ = spc

	// #Check token permission to create a new PC.
	permission, err := a.GetPermission(game.PermissionSubset{
		Source: event.Source.String(),
		Target: characterKey,
	})
	if err != nil {
		return err
	}
	if permission.Value <= 0 {
		return game.ErrInvalidAction
	}

	// #Decrease token permission to create a new PC by 1.
	if err := a.SetPermission(game.Permission{
		Source: event.Source.String(),
		Target: characterKey,
		Value:  permission.Value - 1,
	}); err != nil {
		return err
	}

	// #Retrieve template for new PC.
	template, err := a.GetTemplate(game.TemplateSubset{Type: spc.Type.String()})
	if err != nil {
		return err
	}

	// #Create PC from the template.
	pc := PC(template)
	pc.ID = game.NewULID()
	// TODO list of positions config ? Areas config + random ? Define spawn
	pc.Position = game.Vec3{X: rand.Intn(100), rand.Intn(100), rand.Intn(100)}
	if err := a.SetPC(pc); err != nil {
		return err
	}

	// #Add a new listener for PC.
	targetID := ulid.MustParse(h.listeners[rand.Intn(len(h.listeners))])
	listener := game.Listener{ID: pc.ID}
	if err := h.SendListener(listener, targetID); err != nil {
		return err
	}

	return nil
}
