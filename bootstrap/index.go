package bootstrap

import "github.com/gin-gonic/gin"

func BootstrapApp() {
	app := gin.Default()

	route := app

	route.GET("/", func(ctx *gin.Context) {

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
	})
	app.Run(":8000")
}
