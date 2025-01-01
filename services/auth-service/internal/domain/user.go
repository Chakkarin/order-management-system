package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	SaveUser(ctx context.Context, user *User) error

	HasEmail(ctx context.Context, email *string) (*bool, error)
	HasEmailVerified(ctx context.Context, email *string) (*bool, error)
}

type User struct {
	ID        uuid.UUID `gorm:"primaryKey,type:uuid;default:uuid_generate_v4()"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Verified  bool      `gorm:"default:false"`
	Role      string    `gorm:"default:'C'"` //  C = customer, A = admin
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
