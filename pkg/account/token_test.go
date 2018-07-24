package account

import (
	"net"
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalToken(t *testing.T) {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:3400")
	tokens := []Token{
		Token{},
		Token{
			ID:       ulid.NewID(),
			IP:       nil,
			Account:  ulid.NewID(),
			Ping:     2902,
			CorePool: ulid.NewID(),
			SyncPool: ulid.NewID(),
			PC:       ulid.NewID(),
			Entity:   ulid.NewID(),
		},
		Token{
			ID:       ulid.NewID(),
			IP:       addr,
			Account:  ulid.NewID(),
			Ping:     2902,
			CorePool: ulid.NewID(),
			SyncPool: ulid.NewID(),
			PC:       ulid.NewID(),
			Entity:   ulid.NewID(),
		},
		Token{
			ID:       ulid.NewID(),
			IP:       addr,
			Account:  ulid.NewID(),
			Ping:     290223432,
			CorePool: ulid.NewID(),
			SyncPool: ulid.NewID(),
			PC:       ulid.NewID(),
			Entity:   ulid.NewID(),
		},
		Token{
			IP: addr,
		},
	}
	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, token := range tokens {
			raw, err := token.Marshal(nil)
			assert.NoError(t, err)
			var au Token
			_, err = au.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, au, token)
			_, err = au.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, au, token)
		}
	})
	t.Run("unmarshal safe", func(t *testing.T) {
		for _, token := range tokens {
			raw, err := token.Marshal(nil)
			assert.NoError(t, err)
			var au Token
			for i := 0; i < len(raw); i++ {
				_, err = au.UnmarshalSafe(raw[:i])
				_ = err
			}
		}
	})
}
