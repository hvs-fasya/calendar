package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/hvs-fasya/calendar/internal/configure"
)

//SetLoggers set global logger in accordance with environment
func SetLoggers() {
	var globalLogCfg = zap.NewProductionConfig()
	globalLogCfg.DisableCaller = true
	globalLogCfg.DisableStacktrace = true
	globalLogCfg.EncoderConfig.TimeKey = "time"
	globalLogCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	globalLogCfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if configure.Cfg.Env != "prod" {
		globalLogCfg.Encoding = "console"
		globalLogCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}
	var logger, _ = globalLogCfg.Build()

	zap.ReplaceGlobals(logger)
}
