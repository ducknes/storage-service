package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"storage-service/api"
	"storage-service/database"
	"storage-service/database/kafka"
	"storage-service/service"
	"storage-service/settings"
	"storage-service/tools/storagecontext"
	"time"

	"github.com/GOAT-prod/goatlogger"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	mainCtx context.Context
	config  settings.Config
	logger  goatlogger.Logger

	server *http.Server

	storageService service.Storage

	mongo             *mongo.Client
	redis             *redis.Client
	storageRepository database.StorageRepository
	cacheRepository   database.CacheRepository

	kafkaProducer   *kafka.Producer
	kafkaConsumer   *kafka.Consumer
	messageHandeler kafka.MessageHandler
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

	go a.kafkaConsumer.Consume(storagecontext.New(&http.Request{}))
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

	if err := a.kafkaConsumer.Stop(); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось остановить косьюмер: %v", err))
	}

	a.kafkaProducer.Close()
}

func (a *App) initDatabases() {
	a.initMongo()
	a.initRedis()
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

func (a *App) initRedis() {
	redisCtx, cancelFunc := context.WithTimeout(a.mainCtx, 15*time.Second)
	defer cancelFunc()

	redisClient, err := database.NewRedisClient(redisCtx, a.config.Databases.Redis)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось подключиться к redis: %v", err))
		os.Exit(1)
	}

	a.redis = redisClient
}

func (a *App) initRepositories() {
	a.storageRepository = database.NewStorageRepository(a.mongo, a.config.Databases.MongoDB.Database, a.config.Databases.MongoDB.Collection)
	a.cacheRepository = database.NewCacheRepository(a.redis)

	if a.config.Databases.NeedMocks {
		if err := a.storageRepository.TestData(); err != nil {
			a.logger.Error(err.Error())
			return
		}

		a.logger.Info("тестовые данные успешно добавлены")
	}
}

func (a *App) initKafka() {
	producer, err := kafka.NewProducer(a.config.Databases.Kafka.Address, a.config.Databases.Kafka.ProducerTopic)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось инициализировать продюсер: %v", err))
		os.Exit(1)
	}

	a.kafkaProducer = producer
	a.messageHandeler = kafka.NewMessageHandler(a.storageRepository, a.kafkaProducer)

	consumer, err := kafka.NewConsumer(
		a.messageHandeler,
		a.config.Databases.Kafka.Address,
		a.config.Databases.Kafka.ConsumerTopic,
		a.config.Databases.Kafka.ConsumerGroup)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("не удалось инициализировать консюмер: %v", err))
		os.Exit(1)
	}

	a.kafkaConsumer = consumer
}

func (a *App) initServices() {
	a.storageService = service.NewStorageService(a.storageRepository, a.cacheRepository, a.kafkaProducer)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Error("сервер уже запущен")
		os.Exit(1)
	}

	a.server = api.NewServer(a.mainCtx, a.logger, a.config, a.storageService)
}
