package bootstrap

import (
	"Gin_Gorm_Api/configs"
	"Gin_Gorm_Api/database"
	"Gin_Gorm_Api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")

	}

	// init config
	configs.InitConfig()

	// connect to database
	database.ConnectDatabase()

	//intit gin engine
	app := gin.Default()

	//init routes
	routes.InitRoutes(app)

	//run app
	app.Run(":8000")
}
