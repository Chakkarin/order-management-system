package usecases

import (
	"context"
	"order-management-system/services/auth-service/internal/domain"

	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
)

type UserUsecase struct {
	UserRepo domain.UserRepository
	Redis    *redis.Client
	Mq       *amqp.Channel
}

func NewUserUsecase(repo domain.UserRepository, redis *redis.Client, mq *amqp.Channel) UserUsecaseInterface {
	return &UserUsecase{UserRepo: repo, Redis: redis, Mq: mq}
}

type UserUsecaseInterface interface {
	Register(ctx context.Context, user *domain.User) error
	Verify(ctx context.Context, text *string) error

	sendEmailUsecase(email, type_name *string) error
}
