package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type AppDependencies struct {
	PgDB     *gorm.DB
	Redis    *redis.Client
	RabbitMq *amqp.Channel
}
