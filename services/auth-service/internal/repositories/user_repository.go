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

func (r *UserRepository) FindOneByEmail(ctx context.Context, email *string) (*domain.User, error) {
	var user domain.User
	result := r.DB.WithContext(ctx).Where("email = ?", *email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// ไม่พบข้อมูล
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) SaveUser(ctx context.Context, user *domain.User) error {

	existingUser, err := r.FindOneByEmail(ctx, &user.Email)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return r.DB.WithContext(ctx).Create(user).Error
	}

	existingUser.Password = user.Password

	return r.DB.WithContext(ctx).Save(&existingUser).Error
}

func (r *UserRepository) EmailVerified(ctx context.Context, email *string) error {
	return r.DB.WithContext(ctx).Model(&domain.User{}).Where("email = ?", *email).Update("verified", true).Error
}

func (r *UserRepository) HasEmail(ctx context.Context, email *string) (*bool, error) {

	var count int64

	if err := r.DB.WithContext(ctx).Model(&domain.User{}).Where("email = ? ", *email).Limit(1).Count(&count).Error; err != nil {
		return nil, err
	}

	isDuplicate := count > 0
	return &isDuplicate, nil
}

func (r *UserRepository) HasEmailVerified(ctx context.Context, email *string) (*bool, error) {
	var count int64

	if err := r.DB.WithContext(ctx).Model(&domain.User{}).Where("email = ? and verified is true", *email).Limit(1).Count(&count).Error; err != nil {
		return nil, err
	}

	isEmailVerified := count > 0
	return &isEmailVerified, nil
}
