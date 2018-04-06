package scylla

import (
	"strings"

	"github.com/gocql/gocql"
	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/elojah/game_01"
)

// CreateActor is scylla implementation to create Actor.
func (s *Service) CreateActor(actors ...game.Actor) error {
	if len(actors) == 0 {
		return nil
	}
	sActors := actorsNew(append([]game.Actor{}, actors...))
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
func (s *Service) ListActor(subset game.ActorSubset) ([]byte, error) {
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

	b := flatbuffers.NewBuilder(0)
	game.ActorsStart(b)
	game.ActorsStartValVector(b, len(sMap))
	for _, actor := range sMap {

		game.ActorStart(b)
		game.ActorAddToken(b, b.CreateString(actor["uuid"].(string)))
		game.ActorAddHp(b, actor["hp"].(uint32))
		game.ActorAddMp(b, actor["mp"].(uint32))
		game.ActorAddPosition(b, game.CreateVec3(
			b,
			actor["x"].(uint64),
			actor["y"].(uint64),
			actor["z"].(uint64),
		))
		game.ActorsAddVal(b, game.ActorEnd(b))
	}
	game.ActorsEnd(b)
	return b.FinishedBytes(), iter.Close()
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
	for _, u := range actors {
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
		pos := u.Position(nil)
		args = append(args,
			u.Token(),
			u.Hp(),
			u.Mp(),
			pos.X(),
			pos.Y(),
			pos.Z(),
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
	addHp := patch.Addhp(nil)
	if addHp != nil {
		set = append(set, `hp = hp + ?`)
		args = append(args, addHp)
	}
	subHp := patch.Subhp(nil)
	if subHp != nil {
		set = append(set, `hp = hp - ?`)
		args = append(args, subHp)
	}
	addMp := patch.Addmp(nil)
	if addMp != nil {
		set = append(set, `mp = mp + ?`)
		args = append(args, addMp)
	}
	subMp := patch.Submp(nil)
	if subMp != nil {
		set = append(set, `mp = mp - ?`)
		args = append(args, subMp)
	}
	position := patch.Position(nil)
	if position != nil {
		set = append(set, `x = ?`, `y = ?`, `z = ?`)
		args = append(args, position.X(), position.Y(), position.Z())
	}
	if len(set) == 0 {
		return "", []interface{}{}
	}
	return `SET ` + strings.Join(set, ` , `), args
}

func (subset actorSubset) where() (string, []interface{}) {
	var where []string
	var args []interface{}
	for i := 0; i < subset.TokensLength(); i++ {
		if token := subset.Tokens(i); token != nil {
			where = append(where, `uuid IN ? `)
			args = append(args, string(token))
		}
	}
	if len(where) == 0 {
		return "", []interface{}{}
	}
	return `WHERE ` + strings.Join(where, ` AND `), args
}
