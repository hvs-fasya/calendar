package logging

import (
	"go.uber.org/zap"

	"github.com/hvs-fasya/calendar/internal/configure"
)

//SetLoggers set global logger in accordance with environment
func SetLoggers() {
	var logger, _ = zap.NewDevelopment()
	if configure.Cfg.Env == "prod" {
		globalLogCfg := zap.NewProductionConfig()
		globalLogCfg.DisableStacktrace = true
		logger, _ = globalLogCfg.Build()
	}
	zap.ReplaceGlobals(logger)
}
