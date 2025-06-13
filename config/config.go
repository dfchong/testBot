package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	BotToken string
	WebhookURL string
	ListenAddr string
	RedisEnabled bool
	RedisAddr string
	RedisPassword string
	SentryDSN string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	cfg := &Config{
		BotToken: os.Getenv("TOKEN"),
		WebhookURL: os.Getenv("PUBLICURL"),
		ListenAddr: getEnv("LSITEN_ADD", ":3000"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		SentryDSN: os.Getenv("SENTRY_DSN"),
	}

	if cfg.BotToken == "" || cfg.WebhookURL == ""{
		log.Fatal("bot_token and webhook_url is must need ")
	}

	cfg.RedisEnabled = cfg.RedisAddr != ""

	return cfg
}

func getEnv(key, fallback string) string {
	
}