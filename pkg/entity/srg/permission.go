package srg

import (
	"strconv"
	"strings"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/errors"
)

const (
	permissionKey = "eperm:"
)

// GetPermission implemented with redis.
func (s *Store) GetPermission(source string, target string) (entity.Permission, error) {
	val, err := s.Get(permissionKey + source + ":" + target).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Permission{}, err
		}
		return entity.Permission{}, errors.ErrNotFound
	}

	permission := entity.Permission{
		Source: source,
		Target: target,
	}
	value, err := strconv.Atoi(val)
	permission.Value = value
	return permission, err
}

// SetPermission implemented with redis.
func (s *Store) SetPermission(permission entity.Permission) error {
	return s.Set(permissionKey+permission.Source+":"+permission.Target, permission.Value, 0).Err()
}

// DelPermission implemented with redis.
func (s *Store) DelPermission(source string, target string) error {
	return s.Del(permissionKey + source + ":" + target).Err()
}

// ListPermission list all entity permissions of a source.
func (s *Store) ListPermission(source string) ([]entity.Permission, error) {
	vals, err := s.Keys(permissionKey + source + ":*").Result()
	if err != nil {
		return nil, err
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
