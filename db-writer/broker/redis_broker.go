package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisBroker struct {
	Client *redis.Client
}

func NewRedisClient(host, port, username, password string) (*RedisBroker, error) {
	rdb := redis.NewClient(&redis.Options{
		Password: "",
		Addr:     fmt.Sprint(host, ":", port),
		DB:       0,
		//Username: username,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return &RedisBroker{
		Client: rdb,
	}, nil
}

func (b *RedisBroker) Subscriber(ctx context.Context, channel string, onMessageRecieved func(string)) error {
	subscriber := b.Client.Subscribe(ctx, channel)

	for {
		msg, err := subscriber.ReceiveMessage(ctx)

		if err != nil {
			return err
		}
		go func(payload string) {
			// catching paink
			defer func() {
				if err := recover(); err != nil {
					// do something with error, for now print it to the terminal
					fmt.Println("Error: ", err)
				}
			}()
			onMessageRecieved(payload)
		}(msg.Payload)

	}

}
