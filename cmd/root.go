package cmd

import (
	"fmt"
	"os"

	"github.com/mayudev/animethemes-cli/player"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:     "animethemes",
		Short:   "A CLI for animethemes.moe",
		Version: "0.1.0",
		Long: `A command line interface for animethemes.moe,
a simple and consistent repository of anime opening and ending themes. 
Please consider supporting their work if you like it.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Show help
			cmd.Help()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Apply custom config file
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ~/.config/animethemes-cli.yml)")
	// Apply config key for player
	rootCmd.PersistentFlags().StringVarP(&player.Player, "player", "p", "mpv", "player command to use")
	viper.BindPFlag("player", rootCmd.PersistentFlags().Lookup("player"))
	viper.SetDefault("player", "mpv")

	rootCmd.AddCommand(animeCmd)
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(configFile)
	} else {
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(configDir)
		viper.AddConfigPath("/etc/animethemes-cli")
		viper.AddConfigPath(".")

		viper.SetConfigName("animethemes-cli")
		viper.SetConfigType("yml")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
}
