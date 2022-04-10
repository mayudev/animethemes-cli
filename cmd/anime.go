package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var animeCmd = &cobra.Command{
	Use:   "anime",
	Short: "Search for anime",
	Long:  `Search for anime`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Find query
		query := strings.Join(args, " ")

		fmt.Println("query", query)
	},
}

func init() {
	animeCmd.Flags().BoolP("openings", "o", false, "show only opening themes")
	animeCmd.Flags().BoolP("endings", "e", false, "show only ending themes")
	animeCmd.Flags().Uint("op", 0, "choose particular opening to play")
	animeCmd.Flags().Uint("ed", 0, "choose particular ending to play")
	animeCmd.Flags().StringP("group", "g", "all", "name or index of group to show")
	animeCmd.Flags().BoolP("first", "1", false, "skip choices and pick the first anime result")
}
