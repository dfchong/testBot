package handler

import "gopkg.in/telebot.v4"

// 定义机器人对不同消息和命令的响应

func RegisterHandlers(b *telebot.Bot){
	b.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send("欢迎使用alive bot")
	})

	b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		text := ctx.Text()
		return ctx.Send("你发送了： ", text)
	})
}