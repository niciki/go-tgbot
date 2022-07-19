package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommandHandler) All(update tgbotapi.Update) {
	url := "https://api.covid19api.com/summary"

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	answ := new(Updates)
	json.Unmarshal([]byte(body), answ)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = fmt.Sprintf("*Данные на %s*:\n\n\t🦠Заболело: %d(*+%d*)\n\t"+
		"💀Умерло: %d(*+%d*)\n\t 💊Выздоровело: %d\n\t",
		time.Now().Format(time.RFC822),
		answ.Global.TotalConfirmed,
		answ.Global.NewConfirmed,
		answ.Global.TotalDeaths,
		answ.Global.NewDeaths,
		answ.Global.TotalConfirmed-answ.Global.TotalDeaths)
	msg.ParseMode = "markdown"
	c.bot.Send(msg)
}
