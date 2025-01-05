package usecases

import (
	"context"
	"services/auth-service/internal/config"
	"services/auth-service/shared/models"

	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
)

type UserUsecase struct {
	UserRepo models.UserRepository
	Redis    *redis.Client
	Mq       *amqp.Channel
}

func NewUserUsecase(repo models.UserRepository, deps *config.AppDependencies) UserUsecaseInterface {
	return &UserUsecase{UserRepo: repo, Redis: deps.Redis, Mq: deps.RabbitMq}
}

type UserUsecaseInterface interface {
	Register(ctx context.Context, user *models.User) error
	Verify(ctx context.Context, text *string) error

	sendEmailUsecase(email, type_name *string) error
}
