package middleware

import (
	"Gin_Gorm_Api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	claimsData, err := utils.DecodeToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "unauthenticated",
		})
		return
	}

	ctx.Set("claimsData", claimsData)
	ctx.Set("user_id", claimsData["id"])
	ctx.Set("email", claimsData["email"])
	ctx.Set("username", claimsData["username"])

	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	if token != "123" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "Token is not valid",
		})
		return
	}
	ctx.Next()
}
