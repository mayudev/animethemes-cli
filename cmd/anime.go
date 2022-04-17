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
	first    bool
	forever  bool
)

var animeCmd = &cobra.Command{
	Use:     "anime",
	Aliases: []string{"a", "q"},
	Short:   "Search for anime",
	Long:    `Search for anime`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Validate flags
		if forever && (op != 0 || ed != 0) {
			pterm.Error.Println("Do not use --forever with --op or --ed.")
			os.Exit(0)
		}

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
	animeCmd.Flags().BoolVarP(&first, "first", "1", false, "skip choices and pick the first anime result")
	animeCmd.Flags().BoolVarP(&forever, "forever", "f", false, "ask for choice again once finished playing")
}

// grabAnime Grab anime themes by slug
func grabAnime(slug string) {
	result := api.GetAnime(slug)

	in := util.NewInterface(util.Flags{
		OnlyOpenings: openings,
		OnlyEndings:  endings,
		OpeningN:     op,
		EndingN:      ed,
		First:        first,
	})

	for {
		in.AskThemes(result)
		if !forever {
			break
		}
	}

}
