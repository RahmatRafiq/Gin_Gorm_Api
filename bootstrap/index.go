package bootstrap

import (
	"Gin_Gorm_Api/database"
	"Gin_Gorm_Api/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	database.ConnectDatabase()
	app := gin.Default()

	routes.InitRoutes(app)

	app.Run(":8000")
}
