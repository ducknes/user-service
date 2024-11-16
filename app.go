package main

import (
	"context"
	"fmt"
	"github.com/GOAT-prod/goatlogger"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
	"user-service/api"
	"user-service/database"
	"user-service/service"
	"user-service/settings"
)

type App struct {
	mainCtx context.Context
	config  settings.Config
	logger  goatlogger.Logger

	server *fiber.App

	mongo          *mongo.Client
	userRepository database.UserRepository
	userService    service.User
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
		if err := a.server.Listen(fmt.Sprintf(":%d", a.config.Port)); err != nil {
			a.logger.Panic(fmt.Sprintf("не удалось запустить сервер: %v", err))
			os.Exit(1)
		}
	}()
}

func (a *App) Stop(ctx context.Context) {
	stopCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := a.server.Shutdown(); err != nil {
		a.logger.Error(fmt.Sprintf("Не удалось остановить сервер: %v", err))
	}

	if err := a.mongo.Disconnect(stopCtx); err != nil {
		a.logger.Error(fmt.Sprintf("не удалось отключиться от монги: %v", err))
	}
}

func (a *App) initDatabases() {
	a.initMongo()
}

func (a *App) initKafka() {}

func (a *App) initMongo() {
	mongoClient, err := database.MongoConnect(a.mainCtx, a.config.Databases.MongoDB.ConnectionString)
	if err != nil {
		a.logger.Panic(fmt.Sprintf("err while connecting to mongo db: %v", err))
		os.Exit(1)
	}

	a.mongo = mongoClient
}

func (a *App) initRepositories() {
	a.userRepository = database.NewUserRepository(a.mongo, a.config.Databases.MongoDB.Database, a.config.Databases.MongoDB.Collection)
}

func (a *App) initServices() {
	a.userService = service.NewUserService(a.userRepository)
}

func (a *App) initServer() {
	if a.server != nil {
		a.logger.Panic("server already initialized")
		os.Exit(1)
	}

	a.server = api.NewServer(a.userService)
}
