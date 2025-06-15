package service

import (
	"testbot/config"
	"testing"
)

func TestNewBotService(t *testing.T){
	cfg := &config.Config{
		BotToken: "addasdf",
		WebhookURL: "https://xx.com",
	}

	_, err := NewBotService(cfg)
	if err == nil {
		t.Errorf("newbotservice must be faile for fake token")
	}
}