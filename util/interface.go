package util

import (
	"os"

	"github.com/mayudev/animethemes-cli/api"
)

type Flags struct {
	OnlyOpenings bool
	OnlyEndings  bool
	OpeningN     uint
	EndingN      uint
	First        bool
}

type CurrentSeason struct {
	Season int
	Year   int
	Page   int
}

// I don't know how to call it, so I'll call it Interface
// (as an Interface to interact with utils)
type Interface struct {
	Player        player
	Flags         Flags
	CurrentSeason CurrentSeason
	Exit          func()
}

type player interface {
	Play(v *api.Video)
}

// NewInterface creates a new working Interface to use utils
func NewInterface(flags Flags) Interface {
	return Interface{
		Player: Real{},
		Flags:  flags,
		CurrentSeason: CurrentSeason{
			Page: 1,
		},
		Exit: func() {
			os.Exit(0)
		},
	}
}
