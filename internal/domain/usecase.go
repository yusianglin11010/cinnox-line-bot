package domain

import "go.uber.org/zap"

type LineBotUseCase interface {
	ReceiveMessage(logger *zap.Logger, user, content, replyToken string) error

	GetMessage(logger *zap.Logger, user string, startTime, endTime int64) error
	PushMessage(logger *zap.Logger, user, content string) error
}
