package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mayudev/animethemes-cli/api"
	"github.com/mayudev/animethemes-cli/util"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	openings bool
	endings  bool
	op       uint
	ed       uint
	group    string
	first    bool
)

var animeCmd = &cobra.Command{
	Use:   "anime",
	Short: "Search for anime",
	Long:  `Search for anime`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Find query
		query := strings.Join(args, " ")

		// Show loading spinner
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Searching..."
		s.Start()

		// Grab results from API
		result := api.SearchAnime(query)

		// Hide spinner
		s.Stop()

		// Show no results message if no results are found
		if len(result.Anime) == 0 {
			pterm.Error.Println("No results found.")
			os.Exit(0)
		}

		// Skip prompt if -1 flag passed or there's only one result
		if len(result.Anime) == 1 || first {
			// Show the title of the anime that was selected
			pterm.Println(pterm.LightGreen("-> " + result.Anime[0].Name))

			grabAnime(result.Anime[0].Slug)
		} else {
			choices := make([]string, len(result.Anime))

			for i, v := range result.Anime {
				choices[i] = v.Name
			}

			resultIndex := util.SimpleSelection("Select", choices)

			grabAnime(result.Anime[resultIndex].Slug)
		}
	},
}

func init() {
	// Initialize local flags
	animeCmd.Flags().BoolVarP(&openings, "openings", "o", false, "show only opening themes")
	animeCmd.Flags().BoolVarP(&endings, "endings", "e", false, "show only ending themes")
	animeCmd.Flags().UintVar(&op, "op", 0, "choose particular opening to play")
	animeCmd.Flags().UintVar(&ed, "ed", 0, "choose particular ending to play")
	animeCmd.Flags().StringVarP(&group, "group", "g", "all", "name or index of group to show")
	animeCmd.Flags().BoolVarP(&first, "first", "1", false, "skip choices and pick the first anime result")
}

// grabAnime Grab anime themes by slug
func grabAnime(slug string) {
	result := api.GetAnime(slug)

	choices := make([]string, len(result.Themes))

	for i, v := range result.Themes {
		choices[i] = v.Slug + " " + pterm.Sprintf(pterm.LightYellow(v.Song.Title))
	}

	resultIndex := util.SimpleSelection("Select theme to play", choices)

	entries := result.Themes[resultIndex].Entries

	util.NewInterface().AskEntries(entries)
}
