package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of app",
	Long:  "Print the version of app",
	Run: func(cmd *cobra.Command, args []string) {
		println("0.0.1")
	},
}
