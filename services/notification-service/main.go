package main

import (
	"github.com/LareinaCornia/notification-service/config"
	"github.com/LareinaCornia/notification-service/mail"
	myredis "github.com/LareinaCornia/notification-service/redis"
	"github.com/go-redis/redis/v8"
)

func main() {
	cfg := config.LoadConfig()

	rdb := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})

	mailer := &mail.Mailer{
		Host:     cfg.SMTPHost,
		Port:     cfg.SMTPPort,
		Username: cfg.SMTPUser,
		Password: cfg.SMTPPass,
		From:     cfg.FromEmail,
		To:       cfg.ToEmail,
	}

	myredis.SubscribeAndNotify(rdb, mailer)
}
