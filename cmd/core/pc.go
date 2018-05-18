package main

import (
	"math/rand"

	"github.com/elojah/game_01"
)

func (a *app) CreatePC(event game.Event) error {

	spc := event.Action.(game.SetPC)

	// #Retrieve token object for accountID.
	token, err := a.GetToken(event.Source)
	if err != nil {
		return err
	}

	// #Check user permission to create a new PC.
	left, err := a.GetPCLeft(game.PCLeftSubset{AccountID: token.Account})
	if err != nil {
		return err
	}
	if left <= 0 {
		return game.ErrInvalidAction
	}

	// #Decrease token permission to create a new PC by 1.
	if err := a.SetPCLeft(left-1, token.Account); err != nil {
		return err
	}

	// #Retrieve template for new PC.
	template, err := a.GetTemplate(game.TemplateSubset{Type: spc.Type.String()})
	if err != nil {
		return err
	}

	// #Create PC from the template.
	pc := game.PC(template)
	pc.ID = game.NewULID()
	// TODO list of positions config ? Areas config + random ? Define spawn
	pc.Position = game.Vec3{X: 100 * rand.Float64(), Y: 100 * rand.Float64(), Z: 100 * rand.Float64()}
	return a.SetPC(pc, token.Account)
}
