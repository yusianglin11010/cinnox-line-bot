package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/yusianglin11010/cinnox-line-bot/internal/config"
	"github.com/yusianglin11010/cinnox-line-bot/internal/database"
	"go.uber.org/zap"
)

func TestSaveMessage(t *testing.T) {
	logger, _ := zap.NewProduction()
	dbConfig := config.NewMongoConfig(logger)
	database.Initialize(dbConfig)
	defer database.Close()

	mongo := database.GetMongo()
	dbRepo := NewMongoRepo(mongo.Client)

	userID := uuid.New().String()

	// create not existed user
	err := dbRepo.SaveMessage(logger, context.Background(), userID, "Hello", 3)
	assert.Nil(t, err)
	// create another record for existed user
	err = dbRepo.SaveMessage(logger, context.Background(), userID, "你好", 5)
	assert.Nil(t, err)

}

func TestGetMessage(t *testing.T) {
	logger, _ := zap.NewProduction()
	dbConfig := config.NewMongoConfig(logger)
	database.Initialize(dbConfig)
	defer database.Close()
	mongo := database.GetMongo()
	dbRepo := NewMongoRepo(mongo.Client)

	// XXX: this create step should be replaced with the original Mongo script
	// instead of calling the existed repo method
	userID := uuid.New().String()
	err := dbRepo.SaveMessage(logger, context.Background(), userID, "Hello", 3)
	assert.Nil(t, err)
	err = dbRepo.SaveMessage(logger, context.Background(), userID, "你好", 5)
	assert.Nil(t, err)

	res, err := dbRepo.GetMessage(logger, context.Background(), userID, 0, 6)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 2, len(res.Messages))
}
