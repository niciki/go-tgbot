package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *CommandHandler) Source(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ParseMode = "html"
	msg.Text = "Более подробные данные вы можете получить, нажав <a href=\"https://coronavirus.jhu.edu\">сюда</a>"
	c.bot.Send(msg)
}
