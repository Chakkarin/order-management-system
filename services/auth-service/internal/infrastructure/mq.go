package infrastructure

import (
	"log"
	"order-management-system/services/auth-service/internal/config"
	"os"

	"github.com/streadway/amqp"
)

func ConnectMQ(conf *config.Mq) *amqp.Channel {

	dsn := os.Getenv("MQ_URL")
	if dsn == "" {
		log.Panic("❌ MQ_URL is not set")
	}

	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Failed to open a channel: %v", err)
	}

	createQueueName(ch)

	log.Println("✅ connected to mq...")

	return ch
}

func createQueueName(ch *amqp.Channel) {

	nameQueues := []string{config.MQ_VERIFIER_TYPE, config.MQ_FORGOT_PASS_TYPE}

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
			log.Fatalf("❌ Failed to declare a queue [%v]: %v", nameQueue, err)
		}
	}
}
