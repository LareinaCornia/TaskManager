package redis

import (
	"context"
	"log"

	"github.com/LareinaCornia/notification-service/mail"
	"github.com/go-redis/redis/v8"
)

func SubscribeAndNotify(rdb *redis.Client, mailer *mail.Mailer) {
	ctx := context.Background()
	pubsub := rdb.Subscribe(ctx, "task-updates")
	ch := pubsub.Channel()

	log.Println("Listening for task updates on Redis channel...")

	for msg := range ch {
		log.Printf("Received message: %s\n", msg.Payload)
		mailer.Send("Task Update", msg.Payload)
	}
}
