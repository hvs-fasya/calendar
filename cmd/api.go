package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/hvs-fasya/calendar/internal/configure"
	"github.com/hvs-fasya/calendar/internal/grpc"
)

var (
	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Launch cal_service API",
		Long:  "Launch cal_service API",
		Run: func(cmd *cobra.Command, args []string) {
			zap.L().Info("grpc SERVER run", zap.String("port", configure.Cfg.Server.Port))
			if e := grpc.Run(configure.Cfg.Server); e != nil {
				zap.L().Fatal("calendar grpc SERVER [RUN] error: " + e.Error())
			}
		},
	}
)
