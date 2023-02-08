package repository

import (
	"context"

	"github.com/yusianglin11010/cinnox-line-bot/internal/database/model"
	"github.com/yusianglin11010/cinnox-line-bot/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type DBRepo interface {
	SaveMessage(logger *zap.Logger, ctx context.Context, user, message string, messageTime int64) error
	GetMessage(logger *zap.Logger, ctx context.Context, user string, startDate, endDate int64) (*model.LineDocument, error)
}

type mongoRepo struct {
	client *mongo.Client
}

func NewMongoRepo(client *mongo.Client) DBRepo {
	return &mongoRepo{
		client: client,
	}
}

func (m *mongoRepo) SaveMessage(logger *zap.Logger, ctx context.Context, user, message string, messageTime int64) error {

	msg := model.Message{
		Content: message,
		Time:    messageTime,
	}

	col := m.client.Database("message").Collection("line")
	filter := bson.D{{"user", user}}

	// check if user existed
	if err := col.FindOne(ctx, filter).Decode(bson.M{}); err != nil {
		// create a new document if user not exist
		if err == mongo.ErrNoDocuments {
			lineDocument := model.LineDocument{
				User:     user,
				Messages: []model.Message{msg},
			}
			_, err := col.InsertOne(ctx, lineDocument)
			if err != nil {
				logger.Error("insert message fail", zap.Error(err))
				return domain.ErrMongoCreateFail
			} else {
				return nil
			}
		} else {
			logger.Error("unexpected error", zap.Error(err))
			return domain.ErrUnexpected
		}
	}

	// push data to existed user's document
	update := bson.D{{"$push", bson.D{{"messages", msg}}}}
	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("insert message fail", zap.String("user", user), zap.Error(err))
		return domain.ErrMongoCreateFail
	}
	return nil
}

func (m *mongoRepo) GetMessage(logger *zap.Logger, ctx context.Context, user string, startDate, endDate int64) (*model.LineDocument, error) {

	col := m.client.Database("message").Collection("line")

	filter := bson.M{
		"user": user,
		"messages": bson.M{
			"$elemMatch": bson.M{
				"time": bson.M{
					"$gt": startDate,
					"$lt": endDate,
				},
			},
		},
	}

	data := &model.LineDocument{}

	projection := bson.M{"_id": 0}
	err := col.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(data)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrUserNotExisted
		} else {
			logger.Error("find user message failed", zap.Error(err))
			return nil, domain.ErrMongoGetFail
		}
	}

	return data, nil

}
