package cmd

import "github.com/spf13/cobra"

var notifierCmd = &cobra.Command{
	Use:   "notifier",
	Short: "Launch cal_service notifier",
	Long:  "Launch cal_service notifier",
	Run: func(cmd *cobra.Command, args []string) {
		println("notifier")
	},
}
