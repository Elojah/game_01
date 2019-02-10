package svc

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// PermissionService wraps entity permission operations.
type PermissionService struct {
	EntityPermissionStore entity.PermissionStore
}

// CheckSource check if token has owner permission on source.
func (s PermissionService) CheckSource(id gulid.ID, tok gulid.ID) error {

	// #Check permission token/source.
	permission, err := s.EntityPermissionStore.GetPermission(tok.String(), id.String())
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrap(err, "check permission token")
	}
	if err == nil && account.ACL(permission.Value) != account.Owner {
		return errors.Wrap(
			gerrors.ErrInsufficientACLs{
				Value:  permission.Value,
				Source: tok.String(),
				Target: id.String(),
			},
			"check permission token",
		)
	}
	return errors.Wrap(err, "check permission token")
}
