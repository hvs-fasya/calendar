package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(apiCmd)
	RootCmd.AddCommand(schedulerCmd)
	RootCmd.AddCommand(notifierCmd)
	RootCmd.AddCommand(clientCmd)
}

var RootCmd = &cobra.Command{
	Use:   "cal_service",
	Short: "cal_service - runs calendar app microservices",
	Long: "cal_service runs calendar app microservices\n" +
		"process is a single argument - microservice name should be launched (api, scheduler, notifier)\n" +
		"also version and help arguments available",
	Run: func(cmd *cobra.Command, args []string) {},
}
