package usecases

import (
	"context"
	"errors"
	"order-management-system/services/auth-service/internal/domain"
	"order-management-system/services/auth-service/internal/utils"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseInterface interface {
	Register(ctx context.Context, user *domain.User) error
}

type UserUsecase struct {
	UserRepo domain.UserRepository
	Redis    *redis.Client
}

func NewUserUsecase(repo domain.UserRepository, redis *redis.Client) UserUsecaseInterface {
	return &UserUsecase{UserRepo: repo, Redis: redis}
}

func (u *UserUsecase) Register(ctx context.Context, user *domain.User) error {

	// 1. ตรวจสอบ Username และ Email ว่าซ้ำหรือไม่
	if err := u.UserRepo.CheckDuplicate(ctx, user.Username, user.Email); err != nil {
		return errors.New("username or email already exists")
	}

	// 2. Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// 3. บันทึก User ลงฐานข้อมูล
	if err := u.UserRepo.CreateUser(ctx, user); err != nil {
		return errors.New("failed to create user")
	}

	// 4. สร้าง Verification Code และเก็บใน Redis
	verificationCode, err := utils.GenerateVerificationCode()
	if err != nil {
		return errors.New("failed to generate verification code")
	}

	err = u.Redis.Set(ctx, "verification:"+user.Email, verificationCode, 0).Err()
	if err != nil {
		return errors.New("failed to store verification code")
	}

	return nil
}
