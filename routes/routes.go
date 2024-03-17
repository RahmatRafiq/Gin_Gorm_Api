package routes

import (
	"Gin_Gorm_Api/controllers/MahasiswaControllers"
	"Gin_Gorm_Api/controllers/UserController"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	route := app

	route.GET("/user", UserController.GetAllUser)
	route.GET("/mahasiswa", MahasiswaControllers.GetAllMahasiswa)

}