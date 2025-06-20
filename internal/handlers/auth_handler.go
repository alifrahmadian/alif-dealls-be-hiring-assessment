package handlers

import (
	"net/http"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/handlers/dtos"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/handlers/responses"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/services"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/auth"
	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthService services.AuthService
	SecretKey string
	TTL int
}

func NewAuthHandler(authService *services.AuthService, secretKey string, ttl int) *AuthHandler {
	return &AuthHandler{
		AuthService: *authService,
		SecretKey: secretKey,
		TTL: ttl,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dtos.LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Username":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrUsernameRequired.Error())
				return
			case "Password":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrPasswordRequired.Error())
				return
			}
		}
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	userModel, err := h.AuthService.Login(user.Username, user.Password)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := auth.GenerateToken(userModel, h.SecretKey, h.TTL)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SuccessResponse(c, "login successful", token)
}