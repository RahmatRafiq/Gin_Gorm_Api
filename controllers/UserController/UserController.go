package UserController

import (
	"Gin_Gorm_Api/database"
	models "Gin_Gorm_Api/models/UsersModels"
	"net/http"

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

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.Users)
	errDb := database.DB.Table("users").Where("id = ?", id).First(&user).Error
	if errDb != nil || user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Data Found",
		"data":    user,
	})

}

func Create(ctx *gin.Context) {

}
