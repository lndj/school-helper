package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

//The Environment instance
var Environment *EnvConfig

//The Environment variable
type EnvConfig struct {
	//app
	AppEnv  string `env:"APP_ENV,required"`
	AppPort string `env:"APP_PORT"`

	//wechat
	WechatAppID     string `env:"WECHAT_APP_ID,required"`
	WechatAppSecret string `env:"WECHAT_APP_SECRET,required"`
	WechatToken     string `env:"WECHAT_TOKEN,required"`
	WechatAesKey    string `env:"WECHAT_AES_KEY,required"`

	//slack
	SlackBotToken     string `env:"SLACK_BOT_TOKEN"`
	SlackBotId        string `env:"SLACK_BOT_ID"`
	SlackClientSecret string `env:"SLACK_CLIENT_SECRET"`
	SlackVerifyToken  string `env:"SLACK_VERIFY_TOKEN"`
	SlackChannelID    string `env:"SLACK_CHANNEL_ID"`

	// Telegram
	TelegramChatID   int64  `env:"TELEGRAM_CHAT_ID"`
	TelegramIsDebug  bool   `env:"TELEGRAM_IS_DEBUG" envDefault:"false"`
	TelegramApiToken string `env:"TELEGRAM_API_TOKEN"`
}

func init() {
	cfg := new(EnvConfig)
	err := env.Parse(cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	Environment = cfg
}
