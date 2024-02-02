/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
*/
package cmd

import (
	"fmt"

	"conductor/pkg"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "The init function sets up the conductor on a machine",
	Long:  `The init function reads the provided configuration files and sets up the conductor services in OS.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing conductor.")
		// Read the configurations.
		readConfigurations()
		// Call the initializer with the configurations loaded inside the viper module.
		pkg.InitializeCondutor()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
