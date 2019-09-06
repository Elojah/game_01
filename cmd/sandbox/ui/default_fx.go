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
					Frames: continuous(61),
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
					Frames: continuous(61),
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
					Frames: continuous(72),
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
					Frames: continuous(91),
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
					Frames: continuous(61),
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
					Frames: continuous(61),
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
					Frames: continuous(46),
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
					Frames: continuous(86),
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
					Frames: continuous(100),
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
					Frames: continuous(61),
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
					Frames: continuous(61),
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
					Frames: continuous(40),
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
					Frames: continuous(76),
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
					Frames: continuous(61),
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
					Frames: continuous(61),
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
					Frames: continuous(61),
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
					Frames: continuous(61),
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
					Frames: continuous(61),
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
					Frames: continuous(61),
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
					Frames: continuous(31),
					Loop:   true,
				},
			},
			spritesheet: "weaponhit.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
	}
)
