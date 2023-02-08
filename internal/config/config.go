package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type MongoConfig struct {
	User     string
	Password string
	Port     string
}

type LineBotConfig struct {
	Secret string
	Token  string
}

type RestConfig struct {
	Port string
}

func NewMongoConfig(logger *zap.Logger) *MongoConfig {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("read mongo config fail", zap.Error(err))
	}

	user := viper.GetString("mongo.user")
	password := viper.GetString("mongo.password")
	port := viper.GetString("mongo.port")

	return &MongoConfig{
		User:     user,
		Password: password,
		Port:     port,
	}
}

func NewLineBotConfig(logger *zap.Logger) *LineBotConfig {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("read linebot config fail", zap.Error(err))
	}

	secret := viper.GetString("linebot.secret")
	token := viper.GetString("linebot.token")

	return &LineBotConfig{
		Secret: secret,
		Token:  token,
	}
}

func NewRestConfig(logger *zap.Logger) *RestConfig {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("read rest config fail", zap.Error(err))
	}

	port := viper.GetString("rest.port")

	return &RestConfig{
		Port: port,
	}
}
