package controller

import (
	"api-warung-makan/delivery/middleware"
	"api-warung-makan/model"
	"api-warung-makan/usecase"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	rg          *gin.RouterGroup
	authUseCase usecase.AuthUseCase
}

func (cc *AppController) userAuth(ctx *gin.Context) {
	var user model.UserCredential
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(200, gin.H{
			"message": "Can't bind struct",
		})
		return
	}
	token, err := cc.authUseCase.UserAuth(user)
	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func (cc *AppController) getCustomer(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "user",
	})
}

func NewAppController(routerGroup *gin.RouterGroup, authUseCase usecase.AuthUseCase, tokenMdw middleware.AuthTokenMiddleware) *AppController {
	controller := AppController{
		rg:          routerGroup,
		authUseCase: authUseCase,
	}
	controller.rg.POST("/auth", controller.userAuth)
	protectedGroup := controller.rg.Group("/protected", tokenMdw.RequireToken())
	protectedGroup.GET("/user", controller.getCustomer)
	return &controller
}
