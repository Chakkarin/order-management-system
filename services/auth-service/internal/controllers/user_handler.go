package controllers

import (
	"context"
	"log"
	"order-management-system/services/auth-service/internal/domain"
	"order-management-system/services/auth-service/internal/usecases"
	"order-management-system/services/auth-service/internal/utils"
	"order-management-system/services/auth-service/proto/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	log.Println(req.Email, req.Password)

	// Validate request
	if req.Email == "" || req.Password == "" {
		return nil, status.New(codes.InvalidArgument, "username or password not empty").Err()
	}

	if !utils.IsValidEmail(&req.Email) {
		return nil, status.New(codes.InvalidArgument, "invalid email format").Err()
	}

	// Map Request เป็น Domain Model
	user := &domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	// Call Usecase
	if err := h.Usecase.Register(ctx, user); err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	return &auth.RegisterResponse{
		Message: "User registered successfully. Please verify your email.",
	}, nil
}
