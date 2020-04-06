package main

import (
	"github.com/zgfzgf/rabbitmq/mqengine"
	"go.uber.org/zap"
	"os"
)

var config *mqengine.GbeConfig
var logger *zap.Logger

func main() {
	config = mqengine.GetConfig("./receive.json")
	logger = mqengine.GetLog()
	defer logger.Sync()
	StartClient()
	//for i:=0; i<1000000; i++{
	//	logger.Info("log 初始化成功")
	//}
	c := make(chan os.Signal)
	<-c
}
