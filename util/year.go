package util

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mayudev/animethemes-cli/api"
	"github.com/pterm/pterm"
)

var Seasons = []string{"Winter", "Spring", "Summer", "Fall"}

// FetchPage returns results from current page
func (a Interface) FetchCurrentPage() *api.AnimeSearch {
	results := api.GetSeason(
		Seasons[a.CurrentSeason.Season],
		a.CurrentSeason.Year,
		a.CurrentSeason.Page)

	if len(results.Anime) == 0 {
		pterm.Error.Println("No results found.")
		a.Exit()
	}

	return results
}

func (a Interface) AskSeason() {
	// Show spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Loading..."
	s.Start()

	// Query API
	results := a.FetchCurrentPage()

	s.Stop()

	if len(results.Anime) == 0 {
		pterm.Error.Println("No results found.")
		os.Exit(0)
	}

	// Show results to user
	animes := []string{}

	navigationStyle := pterm.NewStyle(pterm.FgLightGreen, pterm.Bold)
	// If it's not the first page, add a "Previous page" choice
	if a.CurrentSeason.Page != 1 {
		animes = append(animes, navigationStyle.Sprint("Previous page"))
	}

	// Append results
	for _, v := range results.Anime {
		animes = append(animes, v.Name)
	}

	// Figure out if there's another page
	if HasNextPage(&results.Links) {
		animes = append(animes, navigationStyle.Sprint("Next page"))
	}

	// Display prompt
	animeIndex := SimpleSelection("Select", animes)

	if animeIndex == 0 && a.CurrentSeason.Page != 1 {
		// "Previous page" option was selected
		a.CurrentSeason.Page--

		a.AskSeason()

	} else if animeIndex == len(animes)-1 && HasNextPage(&results.Links) {
		// "Next page" option was selected

		// Increase current page
		a.CurrentSeason.Page++

		// Ask again
		a.AskSeason()
	} else {
		// An actual entry was selected

		// Compute real index (before Previous page option was added)
		realIndex := animeIndex

		if a.CurrentSeason.Page != 1 {
			realIndex--
		}

		// Show loading spinner
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Loading..."
		s.Start()

		result := api.GetAnime(results.Anime[realIndex].Slug)

		// Hide spinner
		s.Stop()

		a.AskThemes(result)

	}
}

func HasNextPage(links *api.Links) bool {
	return links.Next != ""
}
