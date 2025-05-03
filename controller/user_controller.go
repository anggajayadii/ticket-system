package controller

import (
	"errors"
	"net/http"
	"ticket-system/dto"
	"ticket-system/middleware"
	"ticket-system/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.ErrorResponse(err))
		return
	}

	user, err := c.authService.Register(req)
	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			ctx.JSON(http.StatusConflict, middleware.ErrorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, middleware.ErrorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"data":    user.ToResponse(),
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.ErrorResponse(err))
		return
	}

	user, err := c.authService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, middleware.ErrorResponse(service.ErrInvalidCreds))
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, middleware.ErrorResponse(err))
		return
	}

	response := user.ToResponse()
	response.Token = token

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data":    response,
	})
}
