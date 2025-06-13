package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	tb "gopkg.in/telebot.v4"
)




func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error for loading .env file")
	}

	webhook := &tb.Webhook{
		Listen: ":8080", //此处是重点 只写端口，不写地址
		Endpoint: &tb.WebhookEndpoint{
			PublicURL: os.Getenv("PUBLICURL"),
		},
	}

	pref := tb.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: webhook,
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	//========================== 以上固定内容  不要改 ＝＝＝＝＝＝＝＝＝＝＝＝＝＝

	btn1 := tb.InlineButton{
		Unique: "btn1",
		Text: "按钮1",
	}

	btn2 := tb.InlineButton{
		Unique: "btn2",
		Text: "按钮2",
	}

	inlineKeys := [][]tb.InlineButton{
		{btn1,btn2},
	}

	bot.Handle(&btn1, func(ctx tb.Context) error {
		return ctx.Send("u clike btn1")
	})

	bot.Handle(&btn2, func(ctx tb.Context) error {
		return ctx.Send("u click btn2")
	})
	
	bot.Handle(tb.OnText, func(ctx tb.Context) error {
		return ctx.Send("inlinekey message", &tb.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
	})
	bot.Start()
}
