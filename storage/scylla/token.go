package scylla

import (
	"bytes"
	"net"
	"strings"

	"github.com/gocql/gocql"

	"github.com/elojah/game_01"
)

func (s *Service) ListToken(subset game.TokenSubset) ([]game.Token, error) {
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

	result := make([]game.Token, len(sMap))
	for i, token := range sMap {
		result[i].ID = token["uuid"].([16]byte)
		result[i].Account = token["account"].([16]byte)
		result[i].IP, err = net.ResolveUDPAddr("udp", token["ip"].(string))
		if err != nil {
			return nil, err
		}
		// result[i].Permissions = a
	}
	return result, iter.Close()
}

func (s *Service) AddTokenPermission(subset TokenSubset, permissions game.Permissions) error {
	sSubset := tokenSubset{subset}
	sPatch := tokenPatch{patch}
	update, uArgs := sPatch.update()
	where, wArgs := sSubset.where()
	query := s.Service.Session.Query(update+where, append(uArgs, wArgs...)...)
	return query.Exec()
}

func (s *Service) UpdateTokenPermission(subset TokenSubset, pSubset PermissionSubset, right Right) error {
	sSubset := tokenSubset{subset}
	sPatch := tokenPatch{patch}
	update, uArgs := sPatch.update()
	where, wArgs := sSubset.where()
	query := s.Service.Session.Query(update+where, append(uArgs, wArgs...)...)
	return query.Exec()

}
func (s *Service) DeleteTokenPermission(subset TokenSubset, pSubset PermissionSubset) error {
	sSubset := tokenSubset{subset}
	where, args := sSubset.where()
	query := s.Service.Session.Query(sSubset.delete()+where, args...)
	return query.Exec()

}

type tokenSubset struct {
	game.TokenSubset
}

type permissionsNew game.Permissions

type permissionSubset struct {
	game.PermissionSubset
}

type permissionPatch struct {
	game.PermissionPatch
}

func (subset tokenSubset) sel() string {
	return `
		SELECT
			uuid,
			permissions,
			account,
			ip,
		FROM global.token_
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

func (np permissionsNew) insert() (string, error) {
	var buffer bytes.Buffer
	var where []string
	var args []interface{}
	for key, val := range np {
		where = append(where, `? : ?`)
		args = append(args, key, val)
	}
	if len(where) == 0 {
		return "", []interface{}{}
	}
	buffer.WriteString(`UPDATE global.token_ SET permissions = permissions + `)
	buffer.WriteString(`{`)
	buffer.WriteString(strings.Join(where, `,`))
	buffer.WriteString(`}`)
	return buffer.String(), args
}

func (subset permissionSubset) where() (string, error) {
	var buffer bytes.Buffer
	var where []string
	var args []interface{}
	for key, val := range np {
		where = append(where, `? : ?`)
		args = append(args, key, val)
	}
	if len(where) == 0 {
		return "", []interface{}{}
	}
	buffer.WriteString(`UPDATE global.token_ SET permissions = permissions + `)
	buffer.WriteString(`{`)
	buffer.WriteString(strings.Join(where, `,`))
	buffer.WriteString(`}`)
	return buffer.String(), args
}
