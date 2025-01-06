package repositories

import (
	"context"
	"services/auth-service/shared/models"

	"gorm.io/gorm"
)

type AuthenRepository struct {
	DB *gorm.DB
}

func NewAuthenRepository(db *gorm.DB) *AuthenRepository {
	return &AuthenRepository{DB: db}
}

func (r *AuthenRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *AuthenRepository) FindOneByEmail(ctx context.Context, email *string) (*models.User, error) {
	var user models.User
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

func (r *AuthenRepository) SaveUser(ctx context.Context, user *models.User) error {

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

func (r *AuthenRepository) EmailVerified(ctx context.Context, email *string) error {
	return r.DB.WithContext(ctx).Model(&models.User{}).Where("email = ?", *email).Update("verified", true).Error
}

func (r *AuthenRepository) HasEmail(ctx context.Context, email *string) (*bool, error) {

	var count int64

	if err := r.DB.WithContext(ctx).Model(&models.User{}).Where("email = ? ", *email).Limit(1).Count(&count).Error; err != nil {
		return nil, err
	}

	isDuplicate := count > 0
	return &isDuplicate, nil
}

func (r *AuthenRepository) HasEmailVerified(ctx context.Context, email *string) (*bool, error) {
	var count int64

	if err := r.DB.WithContext(ctx).Model(&models.User{}).Where("email = ? and verified is true", *email).Limit(1).Count(&count).Error; err != nil {
		return nil, err
	}

	isEmailVerified := count > 0
	return &isEmailVerified, nil
}
