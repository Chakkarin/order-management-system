package controllers

import (
	"context"
	"services/auth-service/internal/domains/authen/models"
	"services/auth-service/internal/domains/authen/usecases"
	"services/auth-service/shared/proto/authen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	AuthenGRPCHandlerInterface interface {
		Register(ctx context.Context, req *authen.RegisterRequest) (*authen.RegisterResponse, error)
	}

	AuthenGRPCHandler struct {
		authen.UnimplementedAuthServiceServer
		Usecase usecases.AuthenUsecaseInterface
	}
)

func NewAuthenGRPCHandler(usecase usecases.AuthenUsecaseInterface) *AuthenGRPCHandler {
	return &AuthenGRPCHandler{Usecase: usecase}
}

func (h *AuthenGRPCHandler) Register(ctx context.Context, req *authen.RegisterRequest) (*authen.RegisterResponse, error) {

	// validate
	if err := authValidate(&req.Email, &req.Password); err != nil {
		return nil, status.New(codes.InvalidArgument, err.Error()).Err()
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
