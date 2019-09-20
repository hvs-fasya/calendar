package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/hvs-fasya/calendar/client"
	"github.com/hvs-fasya/calendar/internal/configure"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Launch cal_service Client",
	Long:  "Launch cal_service Client",
	Run: func(cmd *cobra.Command, args []string) {
		zap.L().Info("grpc CLIENT run")
		if e := client.Run(configure.Cfg.Client); e != nil {
			zap.L().Fatal("calendar grpc CLIENT [RUN] error: " + e.Error())
		}
	},
}
