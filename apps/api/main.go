package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/cinnox-line-bot/internal/config"
	"github.com/yusianglin11010/cinnox-line-bot/internal/database"
	"github.com/yusianglin11010/cinnox-line-bot/internal/handler"
	"github.com/yusianglin11010/cinnox-line-bot/internal/middleware"
	"github.com/yusianglin11010/cinnox-line-bot/internal/repository"
	"github.com/yusianglin11010/cinnox-line-bot/internal/repository/linebot"
	"github.com/yusianglin11010/cinnox-line-bot/internal/usecase"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, _ := zap.NewProduction(zap.AddStacktrace(zapcore.FatalLevel))
	dbConfig := config.NewMongoConfig(logger)
	restConfig := config.NewRestConfig(logger)
	lineBotConfig := config.NewLineBotConfig(logger)

	// initialize mongo DB
	database.Initialize(dbConfig)
	defer database.Close()

	dbRepo := repository.NewMongoRepo(database.GetMongo().Client)
	lineBotClient := linebot.NewLineBotClient(lineBotConfig)
	lineBotUseCase := usecase.NewLineBotUseCase(dbRepo, lineBotClient)

	handler := handler.NewLineBotHandler(logger, lineBotUseCase)

	server := gin.New()
	server.Use(cors.Default())
	server.Use(middleware.AddLoggerToContext(logger))
	server.Use(middleware.AddLineBotClient(lineBotClient))
	server.Use(middleware.RecoveryFromPanic())

	server.GET("/alive", handler.GetHealth)
	server.GET("/message", handler.GetMessage)
	server.POST("/message", handler.PushMessage)
	server.POST("/receive/message", handler.ReceiveMessage)

	server.Run(restConfig.Port)

}
