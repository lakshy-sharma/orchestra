/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
*/
package cmd

import (
	"conductor/pkg"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		conductorConfigs := readConfigurations()
		logger := setupLogger()
		logger.Info("Starting conductor")
		pkg.StartConductor(conductorConfigs, logger)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
