package srg

import (
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
)

const (
	permissionKey = "eperm:"
)

// FetchPermission implemented with redis.
func (s *Store) FetchPermission(source string, target string) (entity.Permission, error) {
	val, err := s.Get(permissionKey + source + ":" + target).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Permission{}, errors.Wrapf(err, "fetch permission %s to %s", source, target)
		}
		return entity.Permission{}, errors.Wrapf(
			gerrors.ErrNotFound{Store: permissionKey, Index: source + ":" + target},
			"fetch permission %s to %s",
			source,
			target,
		)
	}

	permission := entity.Permission{
		Source: source,
		Target: target,
	}
	value, err := strconv.Atoi(val)
	permission.Value = value
	return permission, errors.Wrapf(err, "fetch permission %s to %s", source, target)
}

// UpsertPermission implemented with redis.
func (s *Store) UpsertPermission(permission entity.Permission) error {
	return errors.Wrapf(
		s.Set(permissionKey+permission.Source+":"+permission.Target, permission.Value, 0).Err(),
		"upsert permission %s to %s",
		permission.Source,
		permission.Target,
	)
}

// RemovePermission implemented with redis.
func (s *Store) RemovePermission(source string, target string) error {
	return errors.Wrapf(
		s.Del(permissionKey+source+":"+target).Err(),
		"remove permission %s to %s",
		source,
		target,
	)
}

// ListPermission list all entity permissions of a source.
func (s *Store) ListPermission(source string) ([]entity.Permission, error) {
	vals, err := s.Keys(permissionKey + source + ":*").Result()
	if err != nil {
		return nil, errors.Wrapf(err, "list permissions for %s", source)
	}
	permissions := make([]entity.Permission, len(vals))
	for i, val := range vals {
		permissions[i] = entity.Permission{
			Source: source,
			Target: strings.Split(val, ":")[2],
		}
	}
	return permissions, nil
}
