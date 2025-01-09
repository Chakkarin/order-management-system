package mqconn

import (
	"log"
	"services/auth-service/shared/constants"

	"github.com/streadway/amqp"
)

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
