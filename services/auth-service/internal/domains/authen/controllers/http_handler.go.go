package controllers

import (
	"context"
	"net/http"
	"services/auth-service/internal/domains/authen/models"
	"services/auth-service/internal/domains/authen/usecases"
	"services/auth-service/shared/proto/authen"

	"github.com/gin-gonic/gin"
)

type (
	AuthenHTTPHandlerInterface interface {
		Register(ctx context.Context, req *authen.RegisterRequest) (*authen.RegisterResponse, error)
	}

	AuthenHTTPHandler struct {
		usecase usecases.AuthenUsecaseInterface
	}
)

func NewAuthenHTTPHandler(usecase usecases.AuthenUsecaseInterface) *AuthenHTTPHandler {
	return &AuthenHTTPHandler{usecase: usecase}
}

func (h *AuthenHTTPHandler) Register(c *gin.Context) {

	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate
	if err := authValidate(&req.Email, &req.Password); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.usecase.Register(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully. Please verify your email."})

}
