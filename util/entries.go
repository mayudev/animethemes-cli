package util

import (
	"github.com/mayudev/animethemes-cli/api"
)

type Asker struct {
	Player player
}

type player interface {
	Play(v *api.Video)
}

type Real struct{}

func NewAsker() Asker {
	return Asker{
		Player: Real{},
	}
}

func (Real) Play(v *api.Video) {
	PlayVideo(v)
}

func (a Asker) AskEntries(entries []api.AnimeThemeEntry) {
	if len(entries) > 1 {
		choices := make([]string, len(entries))
		for i, v := range entries {
			choices[i] = BuildEntryChoiceString(&v)
		}

		resultIndex := SimpleSelection("Select version to play", choices)

		a.AskVideos(entries[resultIndex].Videos)
	} else {
		a.AskVideos(entries[0].Videos)
	}
}

func (a Asker) AskVideos(videos []api.Video) {
	if len(videos) > 1 {
		choices := make([]string, len(videos))
		for i, v := range videos {
			choices[i] = BuildVideoChoiceString(&v)
		}

		chosenIndex := SimpleSelection("Select quality", choices)
		a.Player.Play(&videos[chosenIndex])
	} else {
		a.Player.Play(&videos[0])
	}
}
