// 定义函数 NewBotService()
//		生成 bot / redis 客户端
//		绑定 bot + handler
//		将 bot 和 redis 客户端 添加到 svc 

// 定义启动 bot 的函数
package service

import (
	"context"
	"log"
	"testbot/config"
	"testbot/handler"

	"github.com/go-redis/redis/v8"
	"gopkg.in/telebot.v4"
)

// svc 包含 bot 和 redis 
type BotService struct{
	Bot *telebot.Bot
	Redis *redis.Client
	Ctx context.Context
}

func NewBotService(cfg *config.Config) (*BotService, error){
	//webhook 域名和端口
	webhook := &telebot.Webhook{
		Listen: cfg.ListenAddr,		// 端口
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: cfg.WebhookURL,		//域名
		},
	}

	// bot 配置参数
	pref := telebot.Settings{
		Token: cfg.BotToken,
		Poller: webhook,
	}

	// 创建 bot
	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, err
	}

	// bot 添加到 svc
	svc := &BotService{
		Bot: b,
		Ctx: context.Background(),
	}

	// bot 与 消息处理器 绑定
	handler.RegisterHandlers(b)

	// redis 客户端创建 rdb
	// rdb 添加到 svc
	if cfg.RedisEnabled{
		rdb := redis.NewClient(&redis.Options{
			Addr: cfg.RedisAddr,
			Password: cfg.RedisPassword,
		})

		if err := rdb.Ping(svc.Ctx).Err(); err != nil{
			log.Fatalf("failed to connect to redis: %v", err)
		}

		svc.Redis = rdb
		log.Println("connected to redis at ", cfg.RedisAddr)
	}

	return svc, nil
}

// 启动 bot
func (s *BotService) Start() error {
	s.Bot.Start()
	return nil
}