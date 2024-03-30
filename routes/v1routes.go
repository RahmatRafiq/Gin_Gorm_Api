package routes

import (
	"Gin_Gorm_Api/controllers/FileControllers"
	"Gin_Gorm_Api/middleware"

	"github.com/gin-gonic/gin"
)

func v1Route(app *gin.RouterGroup) {
	route := app

	authRoute := route.Group("file", middleware.AuthMiddleware)
	authRoute.POST("/file", FileControllers.UploadFile)
	authRoute.DELETE("/file/:filename", FileControllers.HandleDestroyFile)

}
