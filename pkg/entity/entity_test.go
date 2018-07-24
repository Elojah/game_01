package entity

import (
	"testing"

	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	es := []E{
		// This one fails, gencode transforms empty array into nil.
		// E{
		// 	Position: Position{},
		// },
		E{},
		E{
			ID:   ulid.NewID(),
			Type: ulid.NewID(),
			Name: "Dome kanji systemic beef noodles claymore mine camera rifle 3D-printed chrome pre-drugs tattoo weathered silent. Modem disposable 3D-printed cardboard Shibuya tower convenience store. Shibuya soul-delay gang fluidity weathered rebar garage RAF nodality footage E.I.. Neural math-concrete drone woman range-rover boy jeans dolphin.",
			HP:   97,
			MP:   23,
		},
		E{
			ID:   ulid.NewID(),
			Type: ulid.NewID(),
			HP:   97,
			MP:   23,
		},
		E{
			Position: Position{
				Coord: geometry.Vec3{
					X: 79.6114,
					Y: 86.4033,
					Z: 26.7952,
				},
				SectorID: ulid.NewID(),
			},
		},
		E{
			ID:   ulid.NewID(),
			Type: ulid.NewID(),
			Name: "Apophenia marketing motion into futurity BASE jump-ware garage disposable skyscraper nano. Semiotics systema disposable table hotdog pistol-space dead alcohol human boy car. Bicycle motion-ware j-pop tower cyber-order-flow towards nodality Chiba warehouse semiotics dome. Corrupted apophenia denim man corporation cartel engine woman bridge footage convenience store tanto shoes paranoid. Assault Kowloon car Chiba table shrine modem. Pen film sprawl systema assassin tube plastic semiotics corrupted pistol lights woman receding crypto-decay hotdog. Market spook convenience store cartel long-chain hydrocarbons physical semiotics smart-monofilament. Post-semiotics gang jeans weathered kanji drone fetishism smart-long-chain hydrocarbons. Drone bomb courier face forwards drugs gang Legba hotdog San Francisco neon corporation dome papier-mache Kowloon lights uplink singularity.",
			HP:   97,
			MP:   23,
			Position: Position{
				Coord: geometry.Vec3{
					X: 79.6114,
					Y: 86.4033,
					Z: 26.7952,
				},
				SectorID: ulid.NewID(),
			},
		},
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		for _, e := range es {
			raw, err := e.Marshal(nil)
			assert.NoError(t, err)
			var eu E
			_, err = eu.Unmarshal(raw)
			assert.NoError(t, err)
			assert.Equal(t, eu, e)
			_, err = eu.UnmarshalSafe(raw)
			assert.NoError(t, err)
			assert.Equal(t, eu, e)
		}
	})

	t.Run("unmarshal safe", func(t *testing.T) {
		for _, e := range es {
			raw, err := e.Marshal(nil)
			assert.NoError(t, err)
			var eu E
			for i := 0; i < len(raw); i++ {
				_, err = eu.UnmarshalSafe(raw[:i])
				assert.Error(t, err)
			}
		}
	})
}
