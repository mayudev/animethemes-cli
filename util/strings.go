package util

import (
	"strconv"

	"github.com/mayudev/animethemes-cli/api"
	"github.com/pterm/pterm"
)

func BuildThemeString(t *api.AnimeTheme) string {
	result := t.Slug + " " + pterm.Sprintf(pterm.LightYellow(t.Song.Title))

	if len(t.Entries) == 1 && len(t.Entries[0].Videos) == 1 {
		result += " [" + BuildVideoChoiceString(&t.Entries[0].Videos[0]) + "]"
	}

	return result
}

// BuildEntryChoiceString builds a string to be shown to user in selection
// displaying all information about an entry
func BuildEntryChoiceString(e *api.AnimeThemeEntry) string {
	result := ""

	if len(e.Episodes) > 0 {
		result += "[" + pterm.LightBlue(e.Episodes) + "]"
	}

	if e.Nsfw {
		result += " " + pterm.Red("NSFW")
	}

	if e.Spoiler {
		result += " " + pterm.Red("Spoiler")
	}

	if len(e.Videos) == 1 {
		result += " [" + BuildVideoChoiceString(&e.Videos[0]) + "]"
	}

	return result
}

// BuildVideoChoiceString builds a string to be shown to user in selection
// displaying all information about a video
func BuildVideoChoiceString(v *api.Video) string {
	result := ""

	if v.Resolution != 0 {
		result += pterm.LightBlue(strconv.Itoa(v.Resolution) + "p")
	}

	if len(v.Source) > 0 {
		result += " " + pterm.LightMagenta(v.Source)
	}

	if v.NC {
		result += " " + pterm.LightGreen("Creditless")
	}

	if v.Lyrics {
		result += " " + pterm.LightYellow("Lyrics")
	}

	if v.Overlap != "None" {
		result += " " + pterm.LightRed(v.Overlap)
	}

	return result
}
