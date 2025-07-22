//将消息类型与对应的处理器注册到bot

package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/telebot.v4"
)

// 全局变量
// key = secret_token
// value = bot
// 通过 secret_token 查找 bot
var globalBotRegistry = map[string]telebot.Bot{}

// 定义机器人对不同消息和命令的响应

func RegisterHandlers(b *telebot.Bot) {
	b.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send("欢迎使用alive bot")
	})

	b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		text := ctx.Text()
		return ctx.Send("你发送了： ", text)
	})
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {

	//获取请求头中的secret_token
	token := r.Header.Get("X-Telegram-Bot-Api-Secret-Token")

	//检查请求头中是否包含secret_token,如果滑，bot不能使用
	if token == "" {
		http.Error(w, "missing token", http.StatusForbidden)
		return
	}

	//找不到匹配的bot, 不能使用
	bot, ok := globalBotRegistry[token]
	if !ok{
		http.Error(w,"unknow bot", http.StatusForbidden)
		return
	}

	var update telebot.Update

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		
	}
}
