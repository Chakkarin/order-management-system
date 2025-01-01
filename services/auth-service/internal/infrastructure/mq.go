package infrastructure

import (
	"log"
	"order-management-system/services/auth-service/internal/usecases"
	"os"

	"github.com/streadway/amqp"
)

func ConnectMQ() *amqp.Channel {

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

	_, err := ch.QueueDeclare(
		usecases.VERIFIER_TYPE, // ชื่อ Queue
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		log.Fatalf("❌ Failed to declare a queue [%v]: %v", usecases.VERIFIER_TYPE, err)
	}

	_, err = ch.QueueDeclare(
		usecases.FORGOT_PASS_TYPE, // ชื่อ Queue
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		log.Fatalf("❌ Failed to declare a queue [%v]: %v", usecases.FORGOT_PASS_TYPE, err)
	}
}
