package scylla

import (
	"bytes"
	"net"
	"strings"

	"github.com/gocql/gocql"

	"github.com/elojah/game_01"
)

// ListToken is the scylla implementation to list Token.
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

// AddTokenPermission is the scylla implementation to add TokenPermission.
func (s *Service) AddTokenPermission(subset game.TokenSubset, permissions game.Permissions) error {
	sSubset := tokenSubset{subset}
	sPerms := permissionsNew(permissions)

	insert, iArgs := sPerms.insert()
	where, wArgs := sSubset.where()

	query := s.Service.Session.Query(insert+where, append(iArgs, wArgs...)...)
	return query.Exec()
}

// UpdateTokenPermission is the scylla implementation to update TokenPermission.
func (s *Service) UpdateTokenPermission(subset game.TokenSubset, pSubset game.PermissionSubset, patch game.PermissionPatch) error {
	sSubset := tokenSubset{subset}
	spSubset := permissionSubset{pSubset}
	sPatch := permissionPatch{patch}

	update, uArgs := spSubset.update(sPatch)
	where, wArgs := sSubset.where()

	query := s.Service.Session.Query(update+where, append(uArgs, wArgs...)...)
	return query.Exec()

}

// DeleteTokenPermission is the scylla implementation to delete TokenPermission.
func (s *Service) DeleteTokenPermission(subset game.TokenSubset, pSubset game.PermissionSubset) error {
	sSubset := tokenSubset{subset}
	spSubset := permissionSubset{pSubset}

	delete, dArgs := spSubset.delete()
	where, wArgs := sSubset.where()
	query := s.Service.Session.Query(delete+where, append(dArgs, wArgs...)...)
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

func (np permissionsNew) insert() (string, []interface{}) {
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

func (subset permissionSubset) update(patch permissionPatch) (string, []interface{}) {
	var buffer bytes.Buffer
	var where []string
	var args []interface{}
	for _, id := range subset.IDs {
		where = append(where, ` permissions[?] = ? `)
		args = append(args, id, patch.Right)
	}
	if len(where) == 0 {
		return "", []interface{}{}
	}
	buffer.WriteString(`UPDATE global.token_ SET `)
	buffer.WriteString(strings.Join(where, `,`))
	return buffer.String(), args
}

func (subset permissionSubset) delete() (string, []interface{}) {
	var buffer bytes.Buffer
	var where []string
	var args []interface{}
	for _, id := range subset.IDs {
		where = append(where, `?`)
		args = append(args, id)
	}
	if len(where) == 0 {
		return "", []interface{}{}
	}
	buffer.WriteString(`UPDATE global.token_ SET permissions = permissions - `)
	buffer.WriteString(`{`)
	buffer.WriteString(strings.Join(where, `,`))
	buffer.WriteString(`}`)
	return buffer.String(), args
}
