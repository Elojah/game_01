package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/storage"
)

const (
	permissionKey = "permission:"
)

// GetPermission implemented with redis.
func (s *Service) GetPermission(subset entity.PermissionSubset) (entity.Permission, error) {
	val, err := s.Get(permissionKey + subset.Source + ":" + subset.Target).Result()
	if err != nil {
		if err != redis.Nil {
			return entity.Permission{}, err
		}
		return entity.Permission{}, storage.ErrNotFound
	}

	permission := entity.Permission{
		Source: subset.Source,
		Target: subset.Target,
	}
	value, err := strconv.Atoi(val)
	permission.Value = value
	return permission, err
}

// SetPermission implemented with redis.
func (s *Service) SetPermission(permission entity.Permission) error {
	return s.Set(permissionKey+permission.Source+":"+permission.Target, permission.Value, 0).Err()
}

// DelPermission implemented with redis.
func (s *Service) DelPermission(subset entity.PermissionSubset) error {
	return s.Del(permissionKey + subset.Source + ":" + subset.Target).Err()
}
