package main

import (
	"context"
	"github.com/zgfzgf/rabbitmq/mqengine"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

var config *mqengine.GbeConfig
var logger *zap.Logger

func main() {
	config = mqengine.GetConfig("./receive.json")
	logger = mqengine.GetLog()
	defer logger.Sync()
	ctx, cancel := context.WithCancel(context.Background())
	StartClient(ctx)

	c := make(chan os.Signal)
	signal.Notify(c)
	<-c
	cancel()
}
