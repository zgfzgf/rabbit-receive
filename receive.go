package main

import (
	"github.com/streadway/amqp"
	"github.com/zgfzgf/rabbitmq/mqengine"
	"go.uber.org/zap"
	"time"
)

type Recieve struct {
	productId string
}

func NewRecieve(productId string) *Recieve {
	process := &Recieve{
		productId: productId,
	}
	return process
}

func (p *Recieve) OnProccess(message *mqengine.Message) {
	logger.Info("to do",
		zap.ByteString("body", message.Body))
	message.CorrelationDelivery.Ack(false)
}

type Base struct {
	Status        mqengine.MessageStatus
	ProductId     string
	CorrelationId string
	ReplyTo       string
	MessageId     string
	Time          time.Time
}

type LogEnd struct {
	Base
	Num                 int
	CorrelationDelivery *amqp.Delivery
}
