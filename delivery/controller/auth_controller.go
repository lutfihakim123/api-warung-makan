package controller

import (
	"api-warung-makan/model"
	"api-warung-makan/usecase"
	"api-warung-makan/utils/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router      *gin.Engine
	authUseCase usecase.AuthUseCase
}

func (ac *AuthController) LoginKaryawan(ctx *gin.Context) {
	var user model.Credential
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}
	responseUc, err := ac.authUseCase.LoginKaryawan(user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatus(401)
	} else {
		token, err := auth.GenerateToken(user.Username, "example@email.com")
		if err != nil {
			ctx.AbortWithStatus(401)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
			"data":  responseUc,
		})
	}
}

func NewAuthController(router *gin.Engine, authUseCase usecase.AuthUseCase) *AuthController {
	newAuthController := AuthController{
		router:      router,
		authUseCase: authUseCase,
	}
	router.POST("auth", newAuthController.LoginKaryawan)
	router.Use(auth.AuthTokenMiddleware())
	return &newAuthController
}
