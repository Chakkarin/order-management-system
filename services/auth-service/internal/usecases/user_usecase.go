package usecases

import (
	"context"
	"errors"
	"fmt"
	"order-management-system/services/auth-service/internal/domain"

	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseInterface interface {
	Register(ctx context.Context, user *domain.User) error
}

type UserUsecase struct {
	UserRepo domain.UserRepository
	Redis    *redis.Client
	Mq       *amqp.Channel
}

func NewUserUsecase(repo domain.UserRepository, redis *redis.Client, mq *amqp.Channel) UserUsecaseInterface {
	return &UserUsecase{UserRepo: repo, Redis: redis, Mq: mq}
}

func (u *UserUsecase) Register(ctx context.Context, user *domain.User) error {

	// check email exists
	isDupEmail, err := u.UserRepo.HasEmail(ctx, &user.Email)
	if err != nil {
		return err
	}

	saveUser := func() error {

		// Hash Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password : %v", err.Error())
		}
		user.Password = string(hashedPassword)

		// save user
		if err := u.UserRepo.SaveUser(ctx, user); err != nil {
			return err
		}

		// ส่ง email ใหม่
		err = u.sendEmail(&user.Email, &VERIFIER_TYPE)
		if err != nil {
			return err
		}

		return nil
	}

	if *isDupEmail {
		// check email verified
		isEmailVerified, err := u.UserRepo.HasEmailVerified(ctx, &user.Email)
		if err != nil {
			return err
		}

		if !*isEmailVerified {

			if err := saveUser(); err != nil {
				return err
			}

			return nil
		}

		return errors.New("email already exists")
	}

	if err := saveUser(); err != nil {
		return err
	}

	return nil
}
