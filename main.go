package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error for loading .env file")
	}

	pref := telebot.Settings{
		Token: os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello",func(ctx telebot.Context) error {
		return ctx.Send("hellox")
	})

	b.Handle(telebot.OnText,func(ctx telebot.Context) error {
		var (
			user = ctx.Sender()
			text = ctx.Text()
		)
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}

		return ctx.Send(text)
	})

	b.Handle(telebot.OnChannelPost, func(ctx telebot.Context) error {
		msg := ctx.Message()
	})

	b.Handle(telebot.OnPhoto, func(ctx telebot.Context) error {
		photo := ctx.Message().Photo
	})

	b.Handle(telebot.OnQuery,func(ctx telebot.Context) error {
		return ctx.Answer()
	})

	b.Start()
}