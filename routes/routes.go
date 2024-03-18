package routes

import (
	FakultasControllers "Gin_Gorm_Api/controllers/FakultasController"
	"Gin_Gorm_Api/controllers/MahasiswaControllers"
	"Gin_Gorm_Api/controllers/ProdiControllers"
	"Gin_Gorm_Api/controllers/UserController"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	route := app

	route.GET("/user", UserController.GetAllUser)
	route.POST("/user/create", UserController.Create)

	route.GET("/mahasiswa", MahasiswaControllers.GetAllMahasiswa)
	route.GET("/fakultas", FakultasControllers.GetAllFakultas)
	route.GET("/prodi", ProdiControllers.GetAllProdi)

}
