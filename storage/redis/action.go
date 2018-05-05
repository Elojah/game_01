package redis

import (
	"time"

	"github.com/go-redis/redis"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

const (
	actionKey = "action:"
)

// CreateAction implemented with redis.
func (s *Service) CreateAction(action game.Action, id game.ID, ts time.Time) error {
	var raw []byte
	var err error
	switch action.(type) {
	case game.Listener:
		l := storage.NewListener(action.(game.Listener))
		raw, err = l.Marshal(nil)
	}

	if err != nil {
		return err
	}
	return s.ZAddNX(
		actionKey+id.String(),
		redis.Z{
			Score:  float64(ts.UnixNano()),
			Member: raw,
		},
	).Err()
}
