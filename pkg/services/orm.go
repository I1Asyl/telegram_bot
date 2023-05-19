package services

import (
	"fmt"
	"log"
	"os"

	"github.com/I1Asyl/telegram_bot/models"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type OrmService struct {
	engine *xorm.Engine
}

func NewOrm() *OrmService {
	return &OrmService{SetupOrm()}
}
func (orm OrmService) GetChatUser(chatId int64) (bool, models.User) {
	var connection models.Connection
	ok, err := orm.engine.Where("chat_id = ?", chatId).Get(&connection)
	if !ok || err != nil || connection.NationalId == "" {
		return false, models.User{}
	}
	return orm.GetUser(connection.NationalId)
}

func (orm OrmService) GetUser(nationalId string) (bool, models.User) {
	var user models.User
	ok, err := orm.engine.Where("national_id = ?", nationalId).Get(&user)
	if !ok || err != nil {
		return false, models.User{}
	}
	return true, user
}

func (orm OrmService) SetConnection(nationalId string, chatId int64) error {
	var connection models.Connection
	connection.NationalId = nationalId
	connection.ChatId = chatId
	_, err := orm.engine.Insert(&connection)
	return err
}
func (orm OrmService) UpdateConnection(nationalId string, chatId int64) error {
	var connection models.Connection
	connection.NationalId = nationalId
	connection.ChatId = chatId
	_, err := orm.engine.Cols("national_id").Update(&connection)
	return err
}

func (orm OrmService) SetQuestion(question string, chatId int64) error {
	var connection models.Connection
	_, err := orm.engine.Where("chat_id = ?", chatId).Get(&connection)
	if err != nil {
		return err
	}
	connection.Question = question
	_, err = orm.engine.Cols("question").Update(&connection)
	return err
}

func (orm OrmService) GetQuestion(chatId int) (bool, string) {
	var connection models.Connection
	ok, err := orm.engine.Where("chat_id = ?", chatId).Get(&connection)
	if !ok || err != nil {
		return false, ""
	}
	return true, connection.Question

}

func (orm OrmService) SetUser(user models.User) error {
	_, err := orm.engine.Insert(&user)
	return err
}

// SetupOrm sets up the database connection
func SetupOrm() *xorm.Engine {
	username, password, protocol, address, dbname := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@%v(%v)/%v", username, password, protocol, address, dbname)

	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatalf("Error with a database: %v", err)
	}
	return engine
}
