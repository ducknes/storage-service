package main

import (
	"context"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"github.com/shopspring/decimal"
	"os"
	"os/signal"
	"storage-service/settings"
	"syscall"
	"time"
)

// @title storage-service
// @version 1.0
// @description Сервис управления продукцией на складе для goat-logistics
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization

func main() {
	decimalSettings()

	logger := goatlogger.New(settings.GetAppName())
	logger.SetTag("app")

	mainCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	cfg, err := settings.ReadConfig()
	if err != nil {
		logger.Error(err.Error())
	}

	app := NewApp(mainCtx, cfg, logger)
	app.initDatabases()
	app.initRepositories()
	app.initServices()
	app.initServer()

	app.Start()
	logger.Info(fmt.Sprintf("приложение запушено, порт: %d, конфиг: %s.json", cfg.Port, settings.GetEnv()))

	waitTerminate(mainCtx, app.Stop)

	logger.Info("приложение остановлено")
}

func decimalSettings() {
	decimal.MarshalJSONWithoutQuotes = true
	decimal.DivisionPrecision = 2
}

func waitTerminate(mainCtx context.Context, quitFn func(ctx context.Context)) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	if quitFn == nil {
		return
	}

	quitCtx, cancelQuitCtx := context.WithTimeout(mainCtx, time.Second*15)
	defer cancelQuitCtx()

	quitFn(quitCtx)
}
