package util

import "github.com/mayudev/animethemes-cli/api"

// I don't know how to call it, so I'll call it Interface
// (as an Interface to interact with utils)
type Interface struct {
	Player player
}

type player interface {
	Play(v *api.Video)
}

// NewInterface creates a new working Interface to use utils
func NewInterface() Interface {
	return Interface{
		Player: Real{},
	}
}
