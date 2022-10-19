package middleware

import (
	"api-warung-makan/utils/authenticator"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	accToken authenticator.AccessToken
}

func (a *authTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := authHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer", "", -1)
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		token, err := a.accToken.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		if token != nil {
			ctx.Next()
		} else {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
	}
}

func NewTokenValidator(accToken authenticator.AccessToken) AuthTokenMiddleware {
	return &authTokenMiddleware{
		accToken: accToken,
	}
}
