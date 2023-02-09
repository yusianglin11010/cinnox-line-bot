package linebot

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/yusianglin11010/cinnox-line-bot/internal/config"
)

func NewLineBotClient(cfg *config.LineBotConfig) *linebot.Client {

	bot, err := linebot.New(
		cfg.Secret,
		cfg.Token,
	)
	if err != nil {
		panic(fmt.Sprint("initialize linebot client failed: ", err))
	}
	return bot
}
