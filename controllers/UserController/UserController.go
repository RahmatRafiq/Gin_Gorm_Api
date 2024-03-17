package UserController

import "github.com/gin-gonic/gin"

func GetAllUser(ctx *gin.Context) {

	isValidated := false

	if isValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "Bad Request Something Field not Valid",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})

}
