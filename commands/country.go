package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"unicode"

	"github.com/biter777/countries"

	gt "github.com/bas24/googletranslatefree"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommandHandler) CountryStat(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if len(update.Message.Text) < 9 {
		msg.Text = "Введите команду с именем страны"
		c.bot.Send(msg)
		return
	}
	var CountryNameEng string
	CountryName := update.Message.Text[9:]
	if CountryName[0] > unicode.MaxASCII {
		CountryNameEng, _ = gt.Translate(CountryName, "ru", "en")
	} else {
		CountryNameEng = CountryName
	}
	fmt.Print("\n\n\n", CountryNameEng, "\n\n\n")
	CountryAlpha2 := countries.ByName(CountryNameEng).Alpha2()
	url := "https://api.covid19api.com/summary"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	answ := new(Updates)
	json.Unmarshal([]byte(body), answ)
	var CountryRecord CountryStat
	for _, i := range answ.Countries {
		fmt.Print(i.CountryCode, " ", CountryAlpha2, "\n")
		if i.CountryCode == CountryAlpha2 {
			CountryRecord = i
			break
		}
	}
	if CountryRecord == *new(CountryStat) {
		msg.Text = "Введите корректное"
		c.bot.Send(msg)
	}
	msg.Text = fmt.Sprintf("*%s данные на %s*:\n\n\t🦠Заболело: %d(*+%d*)\n\t"+
		"💀Умерло: %d(*+%d*)\n\t 💊Выздоровело: %d\n\t", countries.ByName(CountryAlpha2).StringRus(),
		time.Now().Format(time.RFC822),
		CountryRecord.TotalConfirmed,
		CountryRecord.NewConfirmed,
		CountryRecord.TotalDeaths,
		CountryRecord.NewDeaths,
		CountryRecord.TotalConfirmed-CountryRecord.TotalDeaths)
	msg.ParseMode = "markdown"
	c.bot.Send(msg)
}
