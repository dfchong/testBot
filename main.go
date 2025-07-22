package main

import (
	"log"
	"testbot/config"
	"testbot/service"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	cfg := config.LoadConfig() 	// 从 配置文件 获取配置

	//创建sentry客户端
	if cfg.SentryDSN != "" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: cfg.SentryDSN,
		})
		if err != nil {
			log.Fatalf("sentry init faled: %v", err)
		}
		defer sentry.Flush(2 * time.Second)
	}

	//创建 botservice
	botSvc, err := service.NewBotService(cfg)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("start failed: %v", err)
	}

	//启动 bot
	log.Println("bot starting ...")
	if err := botSvc.Start(); err != nil {
		sentry.CaptureException(err)
		log.Fatalf("bot panic: %v", err)
	}
}
