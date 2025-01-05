package usecases

import (
	"errors"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func (u *UserUsecase) sendEmailUsecase(email, type_name *string) error {

	// ส่ง mq ให้ notification service ส่ง email
	message := fmt.Sprintf(`{"email": "%v", "type_name": "%v" }`, email, *type_name)

	err := u.Mq.Publish(
		"",         // Default Exchange
		*type_name, // ชื่อ Queue
		false,      // ไม่ต้องการ mandatory delivery
		false,      // ไม่ต้องการ immediate delivery
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return errors.New("failed to publish message")
	}

	log.Printf("✅ Message %v published successfully!", *type_name)

	return nil
}
