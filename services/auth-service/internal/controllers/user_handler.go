package controllers

import (
	"context"
	"fmt"
	"order-management-system/services/auth-service/internal/domain"
	"order-management-system/services/auth-service/internal/usecases"
	"order-management-system/services/auth-service/internal/utils"
	"order-management-system/services/auth-service/proto/auth"
)

type UserHandlerInterface interface {
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
}

type UserHandler struct {
	auth.UnimplementedAuthServiceServer
	Usecase usecases.UserUsecaseInterface
}

func NewUserHandler(usecase usecases.UserUsecaseInterface) *UserHandler {
	return &UserHandler{Usecase: usecase}
}

func (h *UserHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	// Validate request
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, fmt.Errorf("all fields are required")
	}

	if !utils.IsValidEmail(&req.Email) {
		return nil, fmt.Errorf("invalid email format")
	}

	// Map Request เป็น Domain Model
	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	// Call Usecase
	if err := h.Usecase.Register(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to register user: %v", err)
	}

	return &auth.RegisterResponse{
		Message: "User registered successfully. Please verify your email.",
	}, nil
}
