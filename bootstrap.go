package main

import (
	"context"
	"github.com/zgfzgf/rabbitmq/mqengine"
)

func StartClient(ctx context.Context) {
	productId := "aaa"
	readMq := mqengine.NewReaderMQ(productId)
	send := NewClient(productId, readMq)
	send.Start(ctx)
}
