package service

import (
	"context"
	"log"
	"testbot/config"
	"testbot/handler"

	"github.com/go-redis/redis/v8"
	"gopkg.in/telebot.v4"
)


type BotService struct{
	Bot *telebot.Bot
	Redis *redis.Client
	Ctx context.Context
}

func NewBotService(cfg *config.Config) (*BotService, error){
	webhook := &telebot.Webhook{
		Listen: cfg.ListenAddr,
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: cfg.WebhookURL,
		},
	}

	pref := telebot.Settings{
		Token: cfg.BotToken,
		Poller: webhook,
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, err
	}

	svc := &BotService{
		Bot: b,
		Ctx: context.Background(),
	}


	handler.RegisterHandlers(b)

	if cfg.RedisEnabled{
		rdb := redis.NewClient(&redis.Options{
			Addr: cfg.RedisAddr,
			Password: cfg.RedisPassword,
		})

		if err := rdb.Ping(svc.Ctx).Err(); err != nil{
			log.Fatal("failed to connect to redis: %v", err)
		}

		svc.Redis = rdb
		log.Println("connected to redis at ", cfg.RedisAddr)
	}

	return svc, nil
}

func (s *BotService) Start() error {
	s.Bot.Start()
	return nil
}