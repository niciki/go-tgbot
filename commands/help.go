package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *CommandHandler) Help(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Список команд: \n\t*/all* - статистика заражений и смертений по всему миру\n\t" +
		"*/country Имя_страны* - статистика по отдельной стране\n\t" +
		"*/source* - посмотреть источник данных"
	msg.ParseMode = "markdown"
	c.bot.Send(msg)
}
