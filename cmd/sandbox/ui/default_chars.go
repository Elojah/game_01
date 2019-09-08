package ui

import (
	engoc "github.com/EngoEngine/engo/common"
)

var (
	chars = []Entity{
		{
			Player: true,
			SpaceComponent: engoc.SpaceComponent{
				Width:  32,
				Height: 48,
			},
			animations: []*engoc.Animation{
				{
					Name:   "walk_down",
					Frames: continuous(0, 4),
					Loop:   true,
				},
				{
					Name:   "walk_left",
					Frames: continuous(4, 8),
					Loop:   true,
				},
				{
					Name:   "walk_right",
					Frames: continuous(8, 12),
					Loop:   true,
				},
				{
					Name:   "walk_up",
					Frames: continuous(12, 16),
					Loop:   true,
				},
			},
			spritesheet: "whitemage_f.png",
			width:       32,
			height:      48,
			rate:        0.1,
		},
	}
)
