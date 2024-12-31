package repositories

import (
	"context"
	"errors"
	"order-management-system/services/auth-service/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) CheckDuplicate(ctx context.Context, username, email *string) error {

	var count int64

	if err := r.DB.WithContext(ctx).Model(&domain.User{}).Where("username = ?", *username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}

	if err := r.DB.WithContext(ctx).Model(&domain.User{}).Where("email = ?", *email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email already exists")
	}

	return nil
}
