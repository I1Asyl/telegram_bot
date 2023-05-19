package telegram

import (
	"log"

	"github.com/I1Asyl/telegram_bot/pkg/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	*tgbotapi.BotAPI
	services.Services
}

func NewBot(token string, services services.Services) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{bot, services}, nil
}

func (bot *Bot) Start() {
	log.Printf("Authorized on account %s", bot.Self.UserName)
	updates := bot.initUpdates()
	bot.handleUpdates(updates)

}

func (bot *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := bot.handleMessage(update.Message)

			bot.Send(msg)
		}
	}
}

func (bot *Bot) handleMessage(msg *tgbotapi.Message) tgbotapi.MessageConfig {

	_, question := bot.Services.Orm.GetQuestion(int(msg.Chat.ID))
	if question == "national_id" {
		newMsg := tgbotapi.NewMessage(msg.Chat.ID, "Approve your identity by email please")
		err := bot.Services.Orm.SetQuestion("", msg.Chat.ID)
		if err != nil {
			log.Fatal(err)
		}
		err = bot.Services.Orm.UpdateConnection(msg.Text, msg.Chat.ID)
		if err != nil {
			log.Fatal(err)
		}
		return newMsg
	}

	if ok, _ := bot.Services.GetChatUser(msg.Chat.ID); !ok {
		err := bot.Services.Orm.SetConnection("", msg.Chat.ID)
		if err != nil {
			log.Fatal(err)
		}
		newMsg := tgbotapi.NewMessage(msg.Chat.ID, "We do not have your data, please enter your national id")
		err = bot.Services.Orm.SetQuestion("national_id", msg.Chat.ID)
		if err != nil {
			log.Fatal(err)
		}
		return newMsg
	}
	//_, user := bot.Services.GetChatUser(int(msg.Chat.ID))

	newMsg := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	newMsg.ReplyToMessageID = msg.MessageID
	return newMsg
}

func (bot *Bot) initUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return bot.GetUpdatesChan(u)
}
