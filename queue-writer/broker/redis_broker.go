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

func (b *RedisBroker) Publish(ctx context.Context, channel string, payload []byte) error {

	err := b.Client.Publish(ctx, channel, payload).Err()
	if err != nil {
		return err
	}

	return nil
}
