package usecase

import (
	"context"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
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
	if err := uc.dbRepo.SaveMessage(logger, context.TODO(), user, content, time.Now().Unix()); err != nil {
		logger.Error("mongo db save message failed", zap.Error(err))
		return err
	}
	return nil
}

func (uc *lineBotUseCase) GetMessage(logger *zap.Logger, user string, startTime, endTime int64) (*model.LineDocument, error) {

	ctx := context.TODO()
	isUserExist, err := uc.dbRepo.IsUserExist(logger, ctx, user)
	if err != nil {
		return nil, domain.ErrUnexpected
	}
	if !isUserExist {
		return nil, domain.ErrUserNotExisted
	}
	res, err := uc.dbRepo.GetMessage(logger, ctx, user, startTime, endTime)
	if err != nil {
		if err == domain.ErrNoDocuments {
			return nil, err
		}
		logger.Error("mongo db get message failed", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (uc *lineBotUseCase) PushMessage(logger *zap.Logger, user, content string) error {
	ctx := context.TODO()
	isUserExist, err := uc.dbRepo.IsUserExist(logger, ctx, user)

	if err != nil {
		return domain.ErrUnexpected
	}
	if !isUserExist {
		return domain.ErrUserNotExisted
	}

	msg := linebot.NewTextMessage(content)

	_, err = uc.lineBotClient.PushMessage(user, msg).Do()
	if err != nil {
		logger.Error("push message failed", zap.Error(err))
		return err
	}
	return nil
}
