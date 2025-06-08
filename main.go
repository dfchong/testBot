package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("load .env file error")
	}

	webhook := &telebot.Webhook{
		Listen: "127.0.0.1:8080",
		MaxConnections: 100,
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: os.Getenv("PUBLICURL"),
		},
	}

	pref := telebot.Settings{
		Token: os.Getenv("TOKEN"),
		Poller: webhook,
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/test", func(ctx telebot.Context) error {
		return ctx.Send("webhook is working via caddy")
	})

	log.Println("bot running with webhook")
	bot.Start()
}