package mq

import (
	"log"
	"os"
	"services/auth-service/internal/config"
	"services/auth-service/shared/constants"

	"github.com/streadway/amqp"
)

func ConnectMQ(conf *config.Mq) *amqp.Channel {

	dsn := os.Getenv("MQ_URL")
	if dsn == "" {
		log.Panic("❌ MQ_URL is not set")
	}

	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Panicf("❌ Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("❌ Failed to open a channel: %v", err)
	}

	go createQueueName(ch)

	log.Println("✅ connected to mq...")

	return ch
}

func createQueueName(ch *amqp.Channel) {

	nameQueues := []string{
		constants.NAME_VERIFIER_TYPE,
		constants.NAME_FORGOT_PASS_TYPE,
	}

	for _, nameQueue := range nameQueues {
		_, err := ch.QueueDeclare(
			nameQueue, // ชื่อ Queue
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if err != nil {
			log.Panicf("❌ Failed to declare a queue [%v]: %v", nameQueue, err)
		}
	}
}
