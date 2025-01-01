package repositories

import (
	"context"
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

func (r *UserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	return r.DB.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) IsDuplicateUsername(ctx context.Context, username *string) (*bool, error) {

	var count int64

	if err := r.DB.WithContext(ctx).Model(&domain.User{}).Where("username = ? ", *username).Limit(1).Count(&count).Error; err != nil {
		return nil, err
	}

	isDuplicate := count > 0
	return &isDuplicate, nil
}

func (r *UserRepository) IsDuplicateEmail(ctx context.Context, email *string) (*bool, error) {

	var count int64

	if err := r.DB.WithContext(ctx).Model(&domain.User{}).Where("email = ? ", *email).Limit(1).Count(&count).Error; err != nil {
		return nil, err
	}

	isDuplicate := count > 0
	return &isDuplicate, nil
}

func (r *UserRepository) IsEmailVerified(ctx context.Context, email *string) (*bool, error) {
	var count int64

	if err := r.DB.WithContext(ctx).Model(&domain.User{}).Where("email = ? and verified is true", *email).Limit(1).Count(&count).Error; err != nil {
		return nil, err
	}

	isEmailVerified := count > 0
	return &isEmailVerified, nil
}
