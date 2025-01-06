package mq

import (
	"github.com/streadway/amqp"
)

type (
	MQInterface interface {
		PublishingMq(mqConn *amqp.Channel, type_name, message *string) error
	}

	MQ struct {
		Mq *amqp.Channel
	}
)

func NewMQReceiver(mq *amqp.Channel) *MQ {
	return &MQ{Mq: mq}
}

// func PublishingMq(mqConn mq.MQInterface, type_name, message *string) error {

// 	err := mqConn.Publish(
// 		"",         // Default Exchange
// 		*type_name, // ชื่อ Queue
// 		false,      // ไม่ต้องการ mandatory delivery
// 		false,      // ไม่ต้องการ immediate delivery
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        []byte(*message),
// 		},
// 	)
// 	if err != nil {
// 		return errors.New("failed to publish message")
// 	}

// 	log.Printf("✅ Message %v published successfully!", *type_name)

// 	return nil
// }
