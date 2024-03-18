package UserController

import (
	"Gin_Gorm_Api/database"
	models "Gin_Gorm_Api/models/UsersModels"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {

	users := new([]models.Users)

	err := database.DB.Table("users").Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": users,
	})

}

// func CreateUser(ctx *gin.Context) {
// 	var user models.Users
// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.AbortWithStatusJSON(400, gin.H{
// 			"message": "Invalid request body",
// 		})
// 		return
// 	}

// 	if err := database.DB.Table("users").Create(&user).Error; err != nil {
// 		ctx.AbortWithStatusJSON(500, gin.H{
// 			"message": "Internal Server Error",
// 		})
// 		return
// 	}

// 	ctx.JSON(201, gin.H{
// 		"message": "User created successfully",
// 		"data":    user,
// 	})
// }
