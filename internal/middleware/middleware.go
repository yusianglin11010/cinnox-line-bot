package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"
)

type AppContext struct {
	*gin.Context
	logger *zap.Logger

	lineBotClient *linebot.Client
}

func AddLoggerToContext(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("logger", logger)
		ctx.Next()
	}
}

func AddLineBotClient(lineBotClient *linebot.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("lineBotClient", lineBotClient)
		ctx.Next()
	}
}
