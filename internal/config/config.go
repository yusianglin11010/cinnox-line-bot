package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type MongoConfig struct {
	User     string
	Password string
	Port     int
	Host     string
}

type LineBotConfig struct {
	Secret string
	Token  string
}

type RestConfig struct {
	Port string
}

func getConfigPath() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	return path.Join(filepath.Dir(basePath), "..", "config", "config.yaml")
}

func NewMongoConfig(logger *zap.Logger) *MongoConfig {

	cfgPath := getConfigPath()

	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("read mongo config fail", zap.Error(err))
	}

	user := viper.GetString("mongo.user")
	password := viper.GetString("mongo.password")
	port := viper.GetInt("mongo.port")
	host := viper.GetString("mongo.host")

	return &MongoConfig{
		User:     user,
		Password: password,
		Port:     port,
		Host:     host,
	}
}

func NewLineBotConfig(logger *zap.Logger) *LineBotConfig {
	cfgPath := getConfigPath()

	viper.SetConfigFile(cfgPath)
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
	cfgPath := getConfigPath()

	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("read rest config fail", zap.Error(err))
	}

	port := viper.GetString("rest.port")

	return &RestConfig{
		Port: port,
	}
}

func MongoURI(config *MongoConfig) string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)
}

func WriteConfig() error {
	cfgPath := getConfigPath()

	viper.SetConfigFile(cfgPath)
	viper.BindEnv("linebot.secret", "LINE_SECRET")
	viper.BindEnv("linebot.token", "LINE_TOKEN")

	viper.BindEnv("mongo.user", "MONGO_USER")
	viper.BindEnv("mongo.password", "MONGO_PASSWORD")
	viper.BindEnv("mongo.port", "MONGO_PORT")
	viper.BindEnv("mongo.host", "MONGO_HOST")

	viper.BindEnv("rest.port", "REST_PORT")
	err := viper.WriteConfigAs(cfgPath)
	if err != nil {
		return err
	}
	return nil
}
