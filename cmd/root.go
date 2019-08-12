package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/hvs-fasya/calendar/internal/configure"
	"github.com/hvs-fasya/calendar/internal/logging"
	"github.com/hvs-fasya/calendar/internal/storage"
)

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(apiCmd)
	RootCmd.AddCommand(schedulerCmd)
	RootCmd.AddCommand(notifierCmd)
}

var RootCmd = &cobra.Command{
	Use:   "cal_service",
	Short: "cal_service - runs calendar app microservices",
	Long: "cal_service runs calendar app microservices\n" +
		"process is a single argument - microservice name should be launched (api, scheduler, notifier)\n" +
		"also version and help arguments available",
	Run: func(cmd *cobra.Command, args []string) {
		err := configure.LoadConfigs()
		if err != nil {
			fmt.Printf("load configs error: %s\n", err)
			os.Exit(1)
		}
		logging.SetLoggers()
		defer zap.L().Sync()
		zap.L().Info("global logger set")
		storage.Store = storage.InitMemoryStore()
	},
}
