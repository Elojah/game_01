package scylla

import (
	"bytes"
	"strings"

	"github.com/gocql/gocql"

	"github.com/elojah/game_01"
)

func (s *Service) GetToken(Token) (Token, error) {
	sSubset := tokenSubset{subset}
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
	for i, token := range sMap {
		result[i].ID = token["uuid"].([16]byte)
		result[i].HP = token["hp"].(uint8)
		result[i].MP = token["mp"].(uint8)
		result[i].Position.X = token["x"].(float64)
		result[i].Position.Y = token["y"].(float64)
		result[i].Position.Z = token["z"].(float64)
	}
	return result, iter.Close()
}

func (s *Service) AddTokenPermission(TokenSubset, PermissionSubset, Right) error {
	sSubset := tokenSubset{subset}
	sPatch := tokenPatch{patch}
	update, uArgs := sPatch.update()
	where, wArgs := sSubset.where()
	query := s.Service.Session.Query(update+where, append(uArgs, wArgs...)...)
	return query.Exec()
}

func (s *Service) UpdateTokenPermission(PermissionSubset, Right) error {
	sSubset := tokenSubset{subset}
	sPatch := tokenPatch{patch}
	update, uArgs := sPatch.update()
	where, wArgs := sSubset.where()
	query := s.Service.Session.Query(update+where, append(uArgs, wArgs...)...)
	return query.Exec()

}
func (s *Service) DeleteTokenPermission(PermissionSubset) error {
	sSubset := tokenSubset{subset}
	where, args := sSubset.where()
	query := s.Service.Session.Query(sSubset.delete()+where, args...)
	return query.Exec()

}

type permissionSubset struct {
	game.PermissionSubset
}

func (subset tokenSubset) sel() string {
	return `
		SELECT
			uuid,
			permissions,
			account,
			ip,
		FROM global.token
	`
}

func (subset tokenSubset) delete() string {
	return `
		DELETE from global.token
	`
}

func (subset tokenSubset) where() (string, []interface{}) {
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
