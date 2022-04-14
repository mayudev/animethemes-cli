package util

import (
	"os"
	"strconv"

	"github.com/mayudev/animethemes-cli/api"
	"github.com/pterm/pterm"
)

func (a Interface) AskThemes(anime api.Anime) {

	if a.Flags.OpeningN != 0 || a.Flags.EndingN != 0 { // "--op" or "--ed" flag was specified
		sequence := 0
		kind := ""

		if a.Flags.OpeningN != 0 {
			sequence = int(a.Flags.OpeningN)
			kind = "OP"
		} else {
			sequence = int(a.Flags.EndingN)
			kind = "ED"
		}

		for _, v := range anime.Themes {
			if v.Type == kind && v.Sequence == sequence {
				// Match found
				a.AskEntries(v.Entries)
				return
			}
		}

		// No match was found
		pterm.Error.Println(kind + strconv.Itoa(sequence) + " wasn't found.")
		os.Exit(0)

	} else { // Neither flag was specified
		choices := []string{}
		originalIndexes := map[int]int{}

		for i, v := range anime.Themes {
			// Neither flag was specified
			if (!a.Flags.OnlyOpenings && !a.Flags.OnlyEndings) ||
				// Flag "--openings" was specified - allow only openings
				(a.Flags.OnlyOpenings && v.Type == "OP") ||
				// Flag "--endings" was specified - allow only endings
				(a.Flags.OnlyEndings && v.Type == "ED") {
				// Display name
				entry := BuildThemeString(&v)

				// Append to choices slice
				choices = append(choices, entry)

				// Bind filtered index to original index
				originalIndexes[len(choices)-1] = i
			}
		}

		// Check filtered results length
		if len(choices) == 0 {
			pterm.Error.Println("No results found.")
			os.Exit(0)
		}

		// Returns filtered index
		resultIndex := SimpleSelection("Select theme to play", choices)

		// Grab original index
		originalIndex := originalIndexes[resultIndex]

		// Proceed to entry selection
		entries := anime.Themes[originalIndex].Entries
		a.AskEntries(entries)
	}
}

func (a Interface) AskEntries(entries []api.AnimeThemeEntry) {
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

func (a Interface) AskVideos(videos []api.Video) {
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
