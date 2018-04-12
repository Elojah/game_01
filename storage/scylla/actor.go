package scylla

import (
	"bytes"
	"strings"

	"github.com/gocql/gocql"

	"github.com/elojah/game_01"
)

// CreateActor is scylla implementation to create Actor.
func (s *Service) CreateActor(actors []game.Actor) error {
	if len(actors) == 0 {
		return nil
	}
	sActors := actorsNew(actors)
	values, args := sActors.values()
	query := s.Service.Session.Query(sActors.insert()+values, args...)
	return query.Exec()
}

// UpdateActor is scylla implementation to update Actor.
func (s *Service) UpdateActor(subset game.ActorSubset, patch game.ActorPatch) error {
	sSubset := actorSubset{subset}
	sPatch := actorPatch{patch}
	update, uArgs := sPatch.update()
	where, wArgs := sSubset.where()
	query := s.Service.Session.Query(update+where, append(uArgs, wArgs...)...)
	return query.Exec()
}

// DeleteActor is scylla implementation to delete Actor.
func (s *Service) DeleteActor(subset game.ActorSubset) error {
	sSubset := actorSubset{subset}
	where, args := sSubset.where()
	query := s.Service.Session.Query(sSubset.delete()+where, args...)
	return query.Exec()
}

// ListActor is scylla implementation to list Actor.
func (s *Service) ListActor(subset game.ActorSubset) ([]game.Actor, error) {
	sSubset := actorSubset{subset}
	where, args := sSubset.where()
	query := s.Service.Session.
		Query(sSubset.sel()+where, args...).
		Consistency(gocql.One)
	iter := query.Iter()
	sMap, err := iter.SliceMap()
	if err != nil {
		return nil, err
	}

	result := make([]game.Actor, len(sMap))
	for i, actor := range sMap {
		result[i].ID = actor["uuid"].([16]byte)
		result[i].HP = actor["hp"].(uint8)
		result[i].MP = actor["mp"].(uint8)
		result[i].Position.X = actor["x"].(float64)
		result[i].Position.Y = actor["y"].(float64)
		result[i].Position.Z = actor["z"].(float64)
	}
	return result, iter.Close()
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
	var buffer bytes.Buffer
	set, args := patch.set()
	buffer.WriteString(`UPDATE global.actor `)
	buffer.WriteString(set)
	return buffer.String(), args
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
	var buffer bytes.Buffer
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
	buffer.WriteString(`SET `)
	buffer.WriteString(strings.Join(set, ` , `))
	return buffer.String(), args
}

func (subset actorSubset) where() (string, []interface{}) {
	var buffer bytes.Buffer
	var where []string
	var args []interface{}
	for _, id := range subset.IDs {
		where = append(where, `uuid IN ? `)
		args = append(args, string(id[:]))
	}
	if len(where) == 0 {
		return "", []interface{}{}
	}
	buffer.WriteString(`WHERE `)
	buffer.WriteString(strings.Join(where, ` AND `))
	return buffer.String(), args
}
