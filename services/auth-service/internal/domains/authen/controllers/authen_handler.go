package controllers

import (
	"context"
	"log"
	"services/auth-service/internal/domains/authen/usecases"
	"services/auth-service/internal/proto/authen"
	"services/auth-service/shared/models"
	"services/auth-service/shared/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	UserHandlerInterface interface {
		Register(ctx context.Context, req *authen.RegisterRequest) (*authen.RegisterResponse, error)
	}

	UserHandler struct {
		authen.UnimplementedAuthServiceServer
		Usecase usecases.UserUsecaseInterface
	}
)

func NewUserHandler(usecase usecases.UserUsecaseInterface) *UserHandler {
	return &UserHandler{Usecase: usecase}
}

func (h *UserHandler) Register(ctx context.Context, req *authen.RegisterRequest) (*authen.RegisterResponse, error) {

	log.Println(req.Email, req.Password)

	// Validate request
	if req.Email == "" || req.Password == "" {
		return nil, status.New(codes.InvalidArgument, "username or password not empty").Err()
	}

	if !utils.IsEmail(&req.Email) {
		return nil, status.New(codes.InvalidArgument, "invalid email format").Err()
	}

	// Map Request เป็น Domain Model
	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	// Call Usecase
	if err := h.Usecase.Register(ctx, user); err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	return &authen.RegisterResponse{
		Message: "User registered successfully. Please verify your email.",
	}, nil
}
