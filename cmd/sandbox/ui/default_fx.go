package ui

import (
	engoc "github.com/EngoEngine/engo/common"
)

var (
	fxs = []Entity{
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "burn",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "bluefire.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "burn",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "brightfire.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "cast",
					Frames: continuous(0, 72),
					Loop:   true,
				},
			},
			spritesheet: "casting.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "spell",
					Frames: continuous(0, 91),
					Loop:   true,
				},
			},
			spritesheet: "felspell.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "burn",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "fire.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "spin",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "firespin.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "lash",
					Frames: continuous(0, 46),
					Loop:   true,
				},
			},
			spritesheet: "flamelash.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "freeze",
					Frames: continuous(0, 86),
					Loop:   true,
				},
			},
			spritesheet: "freezing.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "load",
					Frames: continuous(0, 100),
					Loop:   true,
				},
			},
			spritesheet: "loading.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "shine",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "magic8.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "pop",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "magicbubbles.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "pop",
					Frames: continuous(0, 40),
					Loop:   true,
				},
			},
			spritesheet: "magickahit.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "spell",
					Frames: continuous(0, 76),
					Loop:   true,
				},
			},
			spritesheet: "magicspell.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "shine",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "midnight.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "explode",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "nebula.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "stand",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "phantom.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "protect",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "protectioncircle.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "burn",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "sunburn.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "burn",
					Frames: continuous(0, 61),
					Loop:   true,
				},
			},
			spritesheet: "vortex.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Width:  100,
				Height: 100,
			},
			animations: []*engoc.Animation{
				{
					Name:   "hit",
					Frames: continuous(0, 31),
					Loop:   true,
				},
			},
			spritesheet: "weaponhit.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
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
