package infrastructure

import (
	"log"
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

	// สร้าง Queue หากยังไม่มี
	_, err = ch.QueueDeclare(
		"notification-queue", // ชื่อ Queue
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalf("❌ Failed to declare a queue: %v", err)
	}

	log.Println("✅ connected to mq...")

	return ch
}
