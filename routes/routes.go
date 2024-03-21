package routes

import (
	"Gin_Gorm_Api/configs/app_config"
	FakultasControllers "Gin_Gorm_Api/controllers/FakultasController"
	"Gin_Gorm_Api/controllers/FileControllers"
	"Gin_Gorm_Api/controllers/MahasiswaControllers"
	"Gin_Gorm_Api/controllers/ProdiControllers"
	"Gin_Gorm_Api/controllers/UserController"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	route := app
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	route.GET("/user", UserController.GetAllUser)
	route.GET("/user/:id", UserController.GetUser)
	route.POST("/user", UserController.Store)
	route.PATCH("/user/:id", UserController.Update)
	route.DELETE("/user/:id", UserController.Destroy)
	route.GET("/user/paginate", UserController.Paginate)
	// /user/paginate?perPage=2 untuk paginasi yang mengambil data 2 per halaman
	// /user/paginate?perPage=3&page=2 untuk paginasi yang mengambil data 3 per halaman dan menampilkan halaman ke 2
	// /user/paginate?page=2 untuk paginasi yang mengambil data 10(default) per halaman dan menampilkan halaman ke 2

	route.POST("/file", FileControllers.UploadFile)
	route.DELETE("/file/:filename", FileControllers.HandleDestroyFile)

	route.GET("/mahasiswa", MahasiswaControllers.GetAllMahasiswa)
	route.GET("/fakultas", FakultasControllers.GetAllFakultas)
	route.GET("/prodi", ProdiControllers.GetAllProdi)

}
