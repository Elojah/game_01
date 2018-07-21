package account

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	as := []A{
		A{},
		A{
			ID:       ulid.NewID(),
			Username: "Zerby",
			Password: "Connell",
			Token:    ulid.NewID(),
		},
		A{
			Username: "",
			Password: "Connell",
		},
		A{
			Username: "Marketing fluidity drone shrine Legba hacker boy youtube. Dissident BASE jump semiotics voodoo god faded boat Legba carbon sub-orbital girl-ware drone j-pop. Bomb fluidity RAF office network franchise warehouse papier-mache pre-apophenia. Garage marketing drone assault nano-skyscraper human silent free-market. Smart-denim uplink rain cardboard lights vinyl DIY soul-delay boat augmented reality long-chain hydrocarbons bicycle. Smart-pistol faded physical motion footage plastic neural. Sub-orbital tower modem nodality A.I. corrupted military-grade otaku shoes. Math-hotdog market RAF tanto vinyl numinous paranoid receding BASE jump neon augmented reality weathered pre-advert meta-j-pop. Nano-receding long-chain hydrocarbons saturation point lights drone-space monofilament meta. ",
			Password: "",
		},
	}
	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, a := range as {
			raw, err := a.Marshal(nil)
			assert.NoError(t, err)
			var au A
			_, err = au.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, au, a)
		}
	})
	t.Run("unmarshal safe", func(t *testing.T) {
		for _, a := range as {
			raw, err := a.Marshal(nil)
			assert.NoError(t, err)
			var au A
			for i := 0; i < len(raw); i++ {
				_, err = au.UnmarshalSafe(raw[:i])
				_ = err
			}
		}
	})
}
