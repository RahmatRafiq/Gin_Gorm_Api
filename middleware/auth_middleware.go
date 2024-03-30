package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if token != "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	ctx.Next()
}
