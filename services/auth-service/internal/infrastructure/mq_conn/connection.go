package mqconn

import (
	"fmt"
	"log"
	"services/auth-service/internal/config"

	"github.com/streadway/amqp"
)

func ConnectMQ(conf *config.Mq) *amqp.Channel {

	dsn := fmt.Sprintf(`amqp://%v:%v@%v:%v/`,
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
	)

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
