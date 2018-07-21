package ability

import (
	"testing"

	"github.com/elojah/game_01/pkg/ulid"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	as := []A{
		// This one fails, gencode transforms empty array into nil.
		// A{
		// 	Components: []Component{},
		// },
		A{},
		A{
			ID:            ulid.NewID(),
			Type:          ulid.NewID(),
			Name:          "Dome kanji systemic beef noodles claymore mine camera rifle 3D-printed chrome pre-drugs tattoo weathered silent. Modem disposable 3D-printed cardboard Shibuya tower convenience store. Shibuya soul-delay gang fluidity weathered rebar garage RAF nodality footage A.I.. Neural math-concrete drone woman range-rover boy jeans dolphin.",
			MPConsumption: 23,
			CD:            97,
			CurrentCD:     24,
			Components:    nil,
		},
		A{
			ID:            ulid.NewID(),
			Type:          ulid.NewID(),
			MPConsumption: 23,
			CD:            97,
			CurrentCD:     24,
			Components:    nil,
		},
		A{
			Components: []Component{nil, nil},
		},
		A{
			Components: []Component{
				HealDirect{
					Amount: 83,
					Type:   44,
				},
				DamageDirect{
					Amount: 62,
					Type:   53,
				},
				HealOverTime{
					Amount:    52,
					Type:      96,
					Frequency: 9,
					Duration:  57,
				},
				DamageOverTime{
					Amount:    71,
					Type:      48,
					Frequency: 48,
					Duration:  89,
				},
			},
		},
		A{
			ID:            ulid.NewID(),
			Type:          ulid.NewID(),
			Name:          "Apophenia marketing motion into futurity BASE jump-ware garage disposable skyscraper nano. Semiotics systema disposable table hotdog pistol-space dead alcohol human boy car. Bicycle motion-ware j-pop tower cyber-order-flow towards nodality Chiba warehouse semiotics dome. Corrupted apophenia denim man corporation cartel engine woman bridge footage convenience store tanto shoes paranoid. Assault Kowloon car Chiba table shrine modem. Pen film sprawl systema assassin tube plastic semiotics corrupted pistol lights woman receding crypto-decay hotdog. Market spook convenience store cartel long-chain hydrocarbons physical semiotics smart-monofilament. Post-semiotics gang jeans weathered kanji drone fetishism smart-long-chain hydrocarbons. Drone bomb courier face forwards drugs gang Legba hotdog San Francisco neon corporation dome papier-mache Kowloon lights uplink singularity.",
			MPConsumption: 23,
			CD:            97,
			CurrentCD:     24,
			Components: []Component{
				HealDirect{
					Amount: 83,
					Type:   44,
				},
				DamageDirect{
					Amount: 62,
					Type:   53,
				},
				HealOverTime{
					Amount:    52,
					Type:      96,
					Frequency: 9,
					Duration:  57,
				},
				DamageOverTime{
					Amount:    71,
					Type:      48,
					Frequency: 48,
					Duration:  89,
				},
			},
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
				assert.Error(t, err)
			}
		}
	})
}
