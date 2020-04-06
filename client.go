package main

import (
	"github.com/zgfzgf/rabbitmq/mqengine"
	"go.uber.org/zap"
)

type Client struct {
	// productId是一个engine的唯一标识，每个product都会对应一个engine
	productId string
	// engine持有的处理
	proccess *Recieve

	// 用于读取消息
	readerHandle *mqengine.RabbitMq
	readerChan   chan *mqengine.Message

}

func NewClient(productId string, reader *mqengine.RabbitMq) *Client {
	e := &Client{
		productId:    productId,
		proccess:     NewRecieve(productId),
		readerHandle: reader,
		readerChan:   make(chan *mqengine.Message, config.ChanNum.Reader),
	}
	return e
}

func (e *Client) Start() {
	if err := recover(); err != nil {
		logger.Error("recover", zap.Error(err.(error)))
	}
	e.readerHandle.RegisterReceiver(e)
	go e.readerHandle.Start()
	go e.runApplier()

}

func (e *Client) runApplier() {
	for {
		select {
		case message := <-e.readerChan:
			e.proccess.OnProccess(message)
		}
	}
}


func (e *Client) Consumer() chan<- *mqengine.Message {
	return e.readerChan
}



