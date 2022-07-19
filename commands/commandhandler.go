package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type GlobalStat struct {
	NewConfirmed   int `json:"NewConfirmed"`
	TotalConfirmed int `json:"TotalConfirmed"`
	NewDeaths      int `json:"NewDeaths"`
	TotalDeaths    int `json:"TotalDeaths"`
	NewRecovered   int `json:"NewRecovered"`
	TotalRecovered int `json:"TotalRecovered"`
}

type CountryStat struct {
	Country        string `json:"Country"`
	CountryCode    string `json:"CountryCode"`
	Slug           string `json:"ala-aland-islands"`
	NewConfirmed   int    `json:"NewConfirmed"`
	TotalConfirmed int    `json:"TotalConfirmed"`
	NewDeaths      int    `json:"NewDeaths"`
	TotalDeaths    int    `json:"TotalDeaths"`
	NewRecovered   int    `json:"NewRecovered"`
	TotalRecovered int    `json:"TotalRecovered"`
	Date           int    `json:"Date"`
}

type Updates struct {
	Global    GlobalStat    `json:"Global"`
	Countries []CountryStat `json:"Countries"`
}

type CommandHandler struct {
	bot *tgbotapi.BotAPI
}

func NewCommandHandler(tgbot *tgbotapi.BotAPI) *CommandHandler {
	return &CommandHandler{bot: tgbot}
}

func (c *CommandHandler) Handle(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			c.Start(update)
		case "help":
			c.Help(update)
		case "all":
			c.All(update)
		case "country":
			c.CountryStat(update)
		case "source":
			c.Source(update)
		default:
			c.Default(update)
		}
	} else {
		c.Default(update)
	}
}
