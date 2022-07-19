package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *CommandHandler) Start(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ParseMode = "html"
	msg.Text = "Это бот по получению статистики по заражению инфекцией covid19, перед началом работы лучше изучить список команд, нажав кнопку /help"
	c.bot.Send(msg)
}
