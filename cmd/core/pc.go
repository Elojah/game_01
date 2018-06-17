package main

import (
	"math/rand"

	"github.com/oklog/ulid"

	"github.com/elojah/game_01"
)

func (a *app) CreatePC(id game.ID, event game.Event) error {

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
	template, err := a.GetEntityTemplate(game.EntityTemplateSubset{Type: spc.Type})
	if err != nil {
		return err
	}

	// #Create PC from the template.
	pc := game.PC(template)
	pc.ID = game.NewID()
	pc.Position = game.Position{
		SectorID: ulid.MustParse("01CF001HTBA3CDR1ERJ6RF183A"),
		Coord:    game.Vec3{X: 100 * rand.Float64(), Y: 100 * rand.Float64(), Z: 100 * rand.Float64()},
	}
	if err := pc.Check(); err != nil {
		return err
	}
	// TODO list of positions config ? Areas config + random ? Define spawn
	return a.SetPC(pc, token.Account)
}

// ConnectPC creates an entity from a PC.
func (a *app) ConnectPC(id game.ID, event game.Event) error {

	cpc := event.Action.(game.ConnectPC)

	// #Retrieve token object for accountID.
	token, err := a.GetToken(event.Source)
	if err != nil {
		return err
	}

	// #Retrieve PC for this account.
	pc, err := a.GetPC(game.PCSubset{
		AccountID: token.Account,
		ID:        cpc.Target,
	})
	if err != nil {
		return err
	}

	entity := game.Entity(pc)
	entity.ID = game.NewID()
	// #Creates entity cloned from pc.
	if err := a.SetEntity(entity, event.TS.UnixNano()); err != nil {
		return err
	}

	// #Creates a new listener for this entity.
	core := a.cores[rand.Intn(len(a.cores))]
	listener := game.Listener{ID: entity.ID}
	return a.SendListener(listener, core)
}
