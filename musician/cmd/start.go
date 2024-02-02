/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
*/
package cmd

import (
	"musician/pkg"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the musician.",
	Long:  `This function starts the musician by setting up the logger and reading the configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		musicianConfigs := readConfigurations()
		logger := setupLogger()
		logger.Info("Starting musician")
		pkg.StartMusician(musicianConfigs, logger)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
