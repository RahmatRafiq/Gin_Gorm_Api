package configs

import (
	"Gin_Gorm_Api/configs/app_config"
	"Gin_Gorm_Api/configs/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}
