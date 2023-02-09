package domain

import (
	"github.com/yusianglin11010/cinnox-line-bot/internal/database/model"
	"go.uber.org/zap"
)

type LineBotUseCase interface {
	ReceiveMessage(logger *zap.Logger, user, content, replyToken string) error

	GetMessage(logger *zap.Logger, user string, startTime, endTime int64) (*model.LineDocument, error)
	PushMessage(logger *zap.Logger, user, content string) error
}
