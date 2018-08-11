package redis

import (
	"strconv"
	"strings"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/storage"
)

const (
	permissionKey = "eperm:"
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

// ListPermission list all entity permissions of a source.
func (s *Service) ListPermission(subset entity.PermissionSubset) ([]entity.Permission, error) {
	vals, err := s.Keys(permissionKey + subset.Source + ":*").Result()
	if err != nil {
		return nil, err
	}
	permissions := make([]entity.Permission, len(vals))
	for i, val := range vals {
		permissions[i] = entity.Permission{
			Source: subset.Source,
			Target: strings.Split(val, ":")[2],
		}
	}
	return permissions, nil
}
