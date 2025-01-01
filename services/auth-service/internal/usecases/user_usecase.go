package usecases

import (
	"context"
	"errors"
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
	isDupEmail, err := u.UserRepo.IsDuplicateEmail(ctx, &user.Email)
	if err != nil {
		return err
	}

	saveUser := func() error {

		// Hash Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("failed to hash password")
		}
		user.Password = string(hashedPassword)

		// save user
		if err := u.UserRepo.SaveUser(ctx, user); err != nil {
			return errors.New("failed to update user")
		}

		// ส่ง email ใหม่
		err = SendVerificationEmail(&user.Email, &VERIFIER_TYPE)
		if err != nil {
			return err
		}

		return nil
	}

	if *isDupEmail {
		// check email verified
		isEmailVerified, err := u.UserRepo.IsEmailVerified(ctx, &user.Email)
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

	// check username exists
	isDupUsername, err := u.UserRepo.IsDuplicateUsername(ctx, &user.Username)
	if err != nil {
		return err
	}

	if *isDupUsername {
		return errors.New("username already exists")
	}

	if err := saveUser(); err != nil {
		return err
	}

	return nil
}
