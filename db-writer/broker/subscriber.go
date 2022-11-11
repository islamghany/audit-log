package broker

import "context"

type Subscriber interface {
	Subscribe(ctx context.Context, channel string, onMessageReciev func(string)) error
}
