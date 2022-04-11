package util

import (
	"strconv"

	"github.com/mayudev/animethemes-cli/api"
	"github.com/mayudev/animethemes-cli/player"
)

func AskEntries(entries []api.AnimeThemeEntry) {
	if len(entries) > 1 {
		choices := make([]string, len(entries))
		for i, v := range entries {
			choices[i] = v.Episodes
		}

		resultIndex := SimpleSelection("Select version to play", choices)

		AskVideos(entries[resultIndex].Videos)
	} else {
		AskVideos(entries[0].Videos)
	}
}

func AskVideos(videos []api.Video) {
	if len(videos) > 1 {
		choices := make([]string, len(videos))
		for i, v := range videos {
			choices[i] = strconv.Itoa(v.Resolution)
		}

		chosenIndex := SimpleSelection("Select quality", choices)
		player.PlayVideo(&videos[chosenIndex])
	} else {
		player.PlayVideo(&videos[0])
	}
}
