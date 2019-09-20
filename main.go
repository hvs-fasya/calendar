package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/hvs-fasya/calendar/cmd"
	"github.com/hvs-fasya/calendar/internal/configure"
	"github.com/hvs-fasya/calendar/internal/logging"
	"github.com/hvs-fasya/calendar/internal/storage"
)

func main() {
	err := configure.LoadConfigs()
	if err != nil {
		fmt.Printf("load configs error: %s\n", err)
		os.Exit(1)
	}

	logging.SetLoggers()
	defer zap.L().Sync()
	zap.L().Info("global logger set")

	storage.Store = storage.InitMemoryStore()

	if err := cmd.RootCmd.Execute(); err != nil {
		zap.L().Fatal(err.Error())
	}
}
