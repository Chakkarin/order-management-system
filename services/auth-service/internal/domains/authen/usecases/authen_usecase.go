package usecases

import (
	"context"
	"errors"
	"fmt"
	"services/auth-service/internal/infrastructure/mq"
	"services/auth-service/shared/models"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecaseInterface interface {
		Register(ctx context.Context, user *models.User) error
		Verify(ctx context.Context, text *string) error

		// sendEmailUsecase(email, type_name *string) error
	}

	AuthenUsecaseDependencies struct {
		Repo     models.AuthenRepositoryInterface
		Redis    *redis.Client
		RabbitMq mq.MQInterface
	}

	AuthenUsecase struct {
		UserRepo models.AuthenRepositoryInterface
		Redis    *redis.Client
		Mq       mq.MQInterface
	}
)

func NewUserUsecase(deps *AuthenUsecaseDependencies) *AuthenUsecase {
	return &AuthenUsecase{UserRepo: deps.Repo, Redis: deps.Redis, Mq: deps.RabbitMq}
}

func (u *AuthenUsecase) Register(ctx context.Context, user *models.User) error {

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
		// message := fmt.Sprintf(`{"email": "%v"}`, user.Email)
		// if err := mq.PublishingMq(u.Mq, &constants.NAME_VERIFIER_TYPE, &message); err != nil {
		// 	return err
		// }

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

func (u *AuthenUsecase) Verify(ctx context.Context, text *string) error {

	/*
		รับ text base64 มา decode email|<string ที่ถูก hash ไว้ใน redis>
		<string ที่ถูก hash ไว้ใน redis> เมื่อถอดออกมาจะได้ email ของ user ที่ต้องการ verify
	*/

	email := text

	// update verified email = true
	return u.UserRepo.EmailVerified(ctx, email)
}
