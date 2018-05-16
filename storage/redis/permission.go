package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	permissionKey = "permission:"
)

// GetPermission implemented with redis.
func (s *Service) GetPermission(subset game.PermissionSubset) (game.Permission, error) {
	val, err := s.Get(permissionKey + subset.Source.String() + ":" + subset.Target.String()).Result()
	if err != nil {
		if err != redis.Nil {
			return game.Permission{}, err
		}
		return game.Permission{}, storage.ErrNotFound
	}

	permission := game.Permission{
		Source: subset.Source,
		Target: subset.Target,
	}
	value, err := strconv.Atoi(val)
	permission.Value = game.Right(value)
	return permission, err
}

// CreatePermission implemented with redis.
func (s *Service) CreatePermission(permission game.Permission) error {
	return s.Set(permissionKey+permission.Source.String()+":"+permission.Target.String(), permission.Value, 0).Err()
}
