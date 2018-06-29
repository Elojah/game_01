package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01/pkg/perm"
	"github.com/elojah/game_01/storage"
)

const (
	permissionKey = "permission:"
)

// GetPermission implemented with redis.
func (s *Service) GetPermission(subset perm.Subset) (perm.P, error) {
	val, err := s.Get(permissionKey + subset.Source + ":" + subset.Target).Result()
	if err != nil {
		if err != redis.Nil {
			return perm.P{}, err
		}
		return perm.P{}, storage.ErrNotFound
	}

	permission := perm.P{
		Source: subset.Source,
		Target: subset.Target,
	}
	value, err := strconv.Atoi(val)
	permission.Value = value
	return permission, err
}

// SetPermission implemented with redis.
func (s *Service) SetPermission(permission perm.P) error {
	return s.Set(permissionKey+permission.Source+":"+permission.Target, permission.Value, 0).Err()
}

// DelPermission implemented with redis.
func (s *Service) DelPermission(subset perm.Subset) error {
	return s.Del(permissionKey + subset.Source + ":" + subset.Target).Err()
}
