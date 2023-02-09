package usecase

import (
	"context"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yusianglin11010/cinnox-line-bot/internal/database/model"
	"github.com/yusianglin11010/cinnox-line-bot/internal/domain"
	"github.com/yusianglin11010/cinnox-line-bot/internal/repository"
	"go.uber.org/zap"
)

type lineBotUseCase struct {
	dbRepo        repository.DBRepo
	lineBotClient *linebot.Client
}

func NewLineBotUseCase(db repository.DBRepo, bot *linebot.Client) domain.LineBotUseCase {
	return &lineBotUseCase{
		dbRepo:        db,
		lineBotClient: bot,
	}
}

func (uc *lineBotUseCase) ReceiveMessage(logger *zap.Logger, user, content, replyToken string) error {
	if _, err := uc.lineBotClient.ReplyMessage(replyToken, linebot.NewTextMessage(domain.ConstReplyMessage)).Do(); err != nil {
		logger.Error("line bot client reply message failed", zap.Error(err))
		return err
	}
	if err := uc.dbRepo.SaveMessage(logger, context.TODO(), user, content, time.Now().UnixMicro()); err != nil {
		logger.Error("mongo db save message failed", zap.Error(err))
		return err
	}
	return nil
}

func (uc *lineBotUseCase) GetMessage(logger *zap.Logger, user string, startTime, endTime int64) (*model.LineDocument, error) {
	res, err := uc.dbRepo.GetMessage(logger, context.TODO(), user, startTime, endTime)
	if err != nil {
		if err == domain.ErrUserNotExisted {
			return nil, err
		}
		logger.Error("mongo db get message failed", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (uc *lineBotUseCase) PushMessage(logger *zap.Logger, user, content string) error {
	panic("not implemented")
}
