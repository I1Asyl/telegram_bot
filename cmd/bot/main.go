package main

import (
	"log"
	"os"

	"github.com/I1Asyl/telegram_bot/pkg/telegram"
	"github.com/joho/godotenv"
)

func init() {
	setupEnv()
}

func main() {
	bot, err := telegram.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	bot.Start()
}

func setupEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}
