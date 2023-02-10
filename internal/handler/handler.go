package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
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

type getMessageReq struct {
	User      string `form:"user" binding:"required"`
	StartTime int64  `form:"start_time"`
	EndTime   int64  `form:"end_time"`
}

type getMessageResp struct {
	Status   string        `json:"status"`
	User     string        `json:"user"`
	Messages []messageResp `json:"messages"`
}

type messageResp struct {
	Time    int64  `json:"time"`
	Content string `json:"content"`
}

func (h *Handler) GetMessage(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)

	req := getMessageReq{}

	if err := c.BindQuery(&req); err != nil {
		logger.Error("bind request failed", zap.Error(err))

		handleError(c, http.StatusBadRequest, domain.ErrInvalidParameter.Error())
		return
	}
	resp := getMessageResp{Status: domain.ConstResponseSuccess, User: req.User}

	lineDocument, err := h.lineBotUseCase.GetMessage(logger, req.User, req.StartTime, req.EndTime)

	if err != nil {
		if err == domain.ErrNoDocuments {
			c.JSON(http.StatusOK, resp)
			return
		} else if err == domain.ErrUserNotExisted {
			handleError(c, http.StatusBadRequest, err.Error())
			return
		} else {
			handleError(c, http.StatusInternalServerError, domain.ErrUnexpected.Error())
			return
		}
	}
	msgResp := []messageResp{}
	for _, msg := range lineDocument.Messages {
		msgResp = append(msgResp,
			messageResp{
				Time:    msg.Time,
				Content: msg.Content,
			})
	}

	resp.Messages = msgResp

	c.JSON(http.StatusOK, resp)
	return

}

type pushMessageReq struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

type pushMessageResp struct {
	Status  string `json:"status"`
	User    string `json:"user"`
	Content string `json:"content"`
}

func (h *Handler) PushMessage(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)

	req := pushMessageReq{}

	if err := c.Bind(&req); err != nil {
		logger.Error("bind request failed", zap.Error(err))

		handleError(c, http.StatusBadRequest, domain.ErrInvalidParameter.Error())
		return
	}

	if err := h.lineBotUseCase.PushMessage(logger, req.User, req.Content); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := pushMessageResp{
		Status:  domain.ConstResponseSuccess,
		User:    req.User,
		Content: req.Content,
	}

	c.JSON(http.StatusOK, resp)
	return
}

type errResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handleError(c *gin.Context, status int, msg string) {

	c.JSON(status, errResp{
		Status:  domain.ConstResponseFailed,
		Message: msg,
	})
}
