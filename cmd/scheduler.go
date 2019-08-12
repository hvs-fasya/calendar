package cmd

import "github.com/spf13/cobra"

var schedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "Launch cal_service scheduler",
	Long:  "Launch cal_service scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		println("scheduler")
	},
}
