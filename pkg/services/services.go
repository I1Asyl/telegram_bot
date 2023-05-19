package services

import "github.com/I1Asyl/telegram_bot/models"

type Services struct {
	Orm
}

type Orm interface {
	SetUser(user models.User) error

	GetChatUser(id int64) (bool, models.User)
	GetUser(nationalId string) (bool, models.User)
	SetQuestion(question string, chatId int64) error
	GetQuestion(chatId int) (bool, string)
	SetConnection(nationalId string, chatId int64) error
	UpdateConnection(nationalId string, chatId int64) error
}

func NewServices() *Services {
	return &Services{
		Orm: NewOrm(),
	}
}
