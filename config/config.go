// 将配置文件中的配置，转到内存struct
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 存储配置
type Config struct {
	BotToken      string
	WebhookURL    string
	ListenAddr    string
	RedisEnabled  bool
	RedisAddr     string
	RedisPassword string
	SentryDSN     string
}

// 将配置从配置文件导入到内存
func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	cfg := &Config{
		BotToken:      os.Getenv("TOKEN"),
		WebhookURL:    os.Getenv("PUBLICURL"),
		ListenAddr:    getEnv("LISTEN_ADDR", ":3000"), //指定监听端口
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		SentryDSN:     os.Getenv("SENTRY_DSN"),
	}

	if cfg.BotToken == "" || cfg.WebhookURL == "" {
		log.Fatal("bot_token and webhook_url is must need ")
	}

	cfg.RedisEnabled = cfg.RedisAddr != ""

	return cfg
}

// 获取webhook publick url, 如果明确给出了addr ,则使用addr , 否则使用“：3000”
// 获取监听端口 default = 3000
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
