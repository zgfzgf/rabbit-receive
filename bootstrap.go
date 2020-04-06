package main

import "github.com/zgfzgf/rabbitmq/mqengine"

func StartClient() {
	productId := "aaa"
	readMq := mqengine.NewReaderMQ(productId)
	send := NewClient(productId, readMq)
	send.Start()
}
