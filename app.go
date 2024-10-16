package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"storage-service/api"
	"storage-service/database"
	"storage-service/service"
	"storage-service/settings"
	"time"
)

type App struct {
	mainCtx context.Context
	config  settings.Config
	logger  goatlogger.Logger

	server *http.Server

	storageService service.Storage

	mongo             *mongo.Client
	storageRepository database.StorageRepository
}

func NewApp(ctx context.Context, config settings.Config, logger goatlogger.Logger) *App {
	return &App{
		mainCtx: ctx,
		config:  config,
		logger:  logger,
	}
}

func (a *App) Start() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error(fmt.Sprintf("приложение неожиданно остановлено, ошибка: %v", err))
		}
	}()
}

func (a *App) Stop(ctx context.Context) {
	stopCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := a.server.Shutdown(stopCtx); err != nil {
		a.logger.Error(fmt.Sprintf("Не удалось остановить сервер: %v", err))
	}

	if err := a.mongo.Disconnect(stopCtx); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось отключиться от монги: %v", err))
	}
}

func (a *App) initDatabases() {
	a.initMongo()
}

func (a *App) initMongo() {
	mongoCtx, cancelFunc := context.WithTimeout(a.mainCtx, 15*time.Second)
	defer cancelFunc()

	mongoClient, err := database.MongoConnect(mongoCtx, a.config.Databases.MongoDB.ConnectionString)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось подключиться к mongoDb, ошибка: %v", err))
		os.Exit(1)
	}

	a.mongo = mongoClient
}

func (a *App) initRepositories() {
	a.storageRepository = database.NewStorageRepository(a.mongo, a.config.Databases.MongoDB.Database, a.config.Databases.MongoDB.Collection)

	if settings.GetEnv() == settings.LocalEnv() && a.config.Databases.NeedMocks {
		if err := a.storageRepository.TestData(); err != nil {
			a.logger.Error(err.Error())
		}
	}
}

func (a *App) initServices() {
	a.storageService = service.NewStorageService(a.storageRepository)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Error("сервер уже запущен")
		os.Exit(1)
	}

	a.server = api.NewServer(a.mainCtx, a.logger, a.config, a.storageService)
}
