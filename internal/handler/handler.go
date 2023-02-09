package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yusianglin11010/cinnox-line-bot/internal/domain"
	"go.uber.org/zap"
)

type Handler struct {
	logger         *zap.Logger
	lineBotUseCase domain.LineBotUseCase
}

func NewLineBotHandler(logger *zap.Logger, uc domain.LineBotUseCase) *Handler {
	return &Handler{
		logger:         logger,
		lineBotUseCase: uc,
	}
}

func (h *Handler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alive",
	})
}

func (h *Handler) ReceiveMessage(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)
	lineBotClient := c.MustGet("lineBotClient").(*linebot.Client)

	events, err := lineBotClient.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			logger.Error("invalid line bot signature", zap.Error(err))
			c.JSON(http.StatusBadRequest, "")
		} else {
			logger.Error("line bot client parse request failed", zap.Error(err))
			c.JSON(http.StatusInternalServerError, domain.ErrUnexpected)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
		}
		message, ok := event.Message.(*linebot.TextMessage)

		// save text message only
		if ok {
			userID := event.Source.UserID
			err := h.lineBotUseCase.ReceiveMessage(logger, userID, message.Text, event.ReplyToken)
			if err != nil {
				c.JSON(http.StatusInternalServerError, domain.ErrUnexpected)
			}
		}
	}
	c.JSON(http.StatusOK, "")
}

func (h *Handler) GetMessage(c *gin.Context) {
	panic("not implemented")
}

func (h *Handler) PushMessage(c *gin.Context) {
	panic("not implemented")
}
