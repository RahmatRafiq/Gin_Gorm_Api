package bootstrap

import (
	"Gin_Gorm_Api/database"
	"Gin_Gorm_Api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")

	}

	database.ConnectDatabase()
	app := gin.Default()

	routes.InitRoutes(app)

	app.Run(":8000")
}
