package alert

import (
	"github.com/lndj/school-helper/config"
	"github.com/lndj/school-helper/utils"
	"gopkg.in/telegram-bot-api.v4"
)

var Bot *tgbotapi.BotAPI

func InitTelegram() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(config.Environment.TelegramApiToken)
	if err != nil {
		utils.Logger.Error("Create telegram client error: ", err)
		return
	}

	Bot.Debug = config.Environment.TelegramIsDebug
}

func SendText(chatID int64, text string) {
	_, err := Bot.Send(tgbotapi.NewMessage(chatID, text))
	if err != nil {
		utils.Logger.Error("Send telegram text msg error. ", err)
	}
}

func SendTextToSchoolHelperBot(text string) {
	SendText(config.Environment.TelegramChatID, text)
}
