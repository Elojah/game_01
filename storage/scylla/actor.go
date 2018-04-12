package scylla

import (
	"strings"

	// "github.com/gocql/gocql"

	"github.com/elojah/game_01"
)

// CreateActor is scylla implementation to create Actor.
func (s *Service) CreateActor(actors ...game.Actor) error {
	return nil
}

// UpdateActor is scylla implementation to update Actor.
func (s *Service) UpdateActor(subset game.ActorSubset, patch game.ActorPatch) error {
	return nil
}

// DeleteActor is scylla implementation to delete Actor.
func (s *Service) DeleteActor(subset game.ActorSubset) error {
	return nil
}

// ListActor is scylla implementation to list Actor.
func (s *Service) ListActor(subset game.ActorSubset) ([]byte, error) {
	return nil, nil
}

type actorsNew []game.Actor
type actorPatch struct {
	game.ActorPatch
}
type actorSubset struct {
	game.ActorSubset
}

func (actors actorsNew) insert() string {
	return `
		INSERT INTO global.actor (
			uuid,
			hp,
			mp,
			x,
			y,
			z
		)
	`
}

func (patch actorPatch) update() (string, []interface{}) {
	set, args := patch.set()
	return `UPDATE global.actor ` + set + ` `, args
}

func (subset actorSubset) sel() string {
	return `
		SELECT
			uuid,
			hp,
			mp,
			x,
			y,
			z
		FROM global.actor
	`
}

func (subset actorSubset) delete() string {
	return `
		DELETE from global.actor
	`
}

func (actors actorsNew) values() (string, []interface{}) {
	var values []string
	var args []interface{}
	for _, actor := range actors {
		values = append(values, `
			(
				?,
				?,
				?,
				?,
				?,
				?
			)
		`)
		args = append(args,
			string(actor.ID[:]),
			actor.HP,
			actor.MP,
			actor.Position.X,
			actor.Position.Y,
			actor.Position.Z,
		)
	}
	if len(values) == 0 {
		return "", []interface{}{}
	}
	return `VALUES ` + strings.Join(values, ` , `), args
}

func (patch actorPatch) set() (string, []interface{}) {
	var set []string
	var args []interface{}
	if patch.HP != nil {
		set = append(set, `hp = ?`)
		args = append(args, patch.HP)
	}
	if patch.MP != nil {
		set = append(set, `mp = ?`)
		args = append(args, patch.MP)
	}
	if patch.Position != nil {
		set = append(set, `x = ?`, `y = ?`, `z = ?`)
		args = append(args, patch.Position.X, patch.Position.Y, patch.Position.Z)
	}
	if len(set) == 0 {
		return "", []interface{}{}
	}
	return `SET ` + strings.Join(set, ` , `), args
}

func (subset actorSubset) where() (string, []interface{}) {
	var where []string
	var args []interface{}
	for _, id := range subset.IDs {
		where = append(where, `uuid IN ? `)
		args = append(args, string(id[:]))
	}
	if len(where) == 0 {
		return "", []interface{}{}
	}
	return `WHERE ` + strings.Join(where, ` AND `), args
}
