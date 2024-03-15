package models

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

import (
	"go-sql-driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/gin_gorm_api"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})
	DB = database, err := 
}
