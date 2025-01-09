package mqconn

import (
	"errors"
	"log"

	"github.com/streadway/amqp"
)

type (
	MqUtilsInterface interface {
		PublishingMq(type_name, message *string) error
	}

	MqUtils struct {
		Mq *amqp.Channel
	}
)

func NewMqUtils(mq *amqp.Channel) *MqUtils {
	return &MqUtils{Mq: mq}
}

func (r *MqUtils) PublishingMq(type_name, message *string) error {

	err := r.Mq.Publish(
		"",         // Default Exchange
		*type_name, // ชื่อ Queue
		false,      // ไม่ต้องการ mandatory delivery
		false,      // ไม่ต้องการ immediate delivery
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(*message),
		},
	)
	if err != nil {
		return errors.New("failed to publish message")
	}

	log.Printf("✅ Message %v published successfully!", *type_name)

	return nil
}
