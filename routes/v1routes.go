package routes

import (
	"Gin_Gorm_Api/controllers/FileControllers"
	"Gin_Gorm_Api/middleware"

	"github.com/gin-gonic/gin"
)

func SendStatus(ctx *gin.Context) {

	filename := ctx.MustGet("filename").(string)

	ctx.JSON(200, gin.H{
		"message": "File Uploaded",
		"name":    filename,
	})
}

func v1Route(app *gin.RouterGroup) {
	route := app

	authRoute := route.Group("file", middleware.AuthMiddleware)
	authRoute.POST("/file", FileControllers.UploadFile)
	authRoute.POST("/middleware", SendStatus)
	authRoute.DELETE("/file/:filename", FileControllers.HandleDestroyFile)

}
