package cmd

import "github.com/spf13/cobra"

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Launch cal_service API",
	Long:  "Launch cal_service API",
	Run: func(cmd *cobra.Command, args []string) {
		println("api")
	},
}
