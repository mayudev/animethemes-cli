package cmd

import (
	"os"
	"strconv"
	"time"

	"github.com/mayudev/animethemes-cli/api"
	"github.com/mayudev/animethemes-cli/util"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var yearCmd = &cobra.Command{
	Use:     "year",
	Aliases: []string{"y"},
	Short:   "Find anime released in a particular year",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Try to parse argument as year
		arg := args[0]
		year, err := strconv.Atoi(arg)

		if err != nil {
			pterm.Error.Println("Incorrect year provided")
			os.Exit(0)
		}

		// Check if year is in animethemes.moe db range
		currentYear := time.Now().Year()

		if year > currentYear || year < 1963 {
			pterm.Error.Printf("Year out of range (1963-%v)\n", currentYear)
			os.Exit(0)
		}

		// Ask about season
		askSeason(year)
	},
}

func askSeason(year int) {
	choices := []string{"Winter", "Spring", "Summer", "Fall"}

	// Show season selection to user
	seasonIndex := util.SimpleSelection("Select season", choices)

	// Query API
	results := api.GetSeason(choices[seasonIndex], year, 1)

	if len(results.Anime) == 0 {
		pterm.Error.Println("No results found.")
		os.Exit(0)
	}

}
