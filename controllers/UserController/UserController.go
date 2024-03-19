package UserController

import (
	"Gin_Gorm_Api/database"
	models "Gin_Gorm_Api/models/UsersModels"
	"Gin_Gorm_Api/requests"
	"Gin_Gorm_Api/responses"
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
	user := new(responses.UserResponses)
	errDb := database.DB.Table("users").Where("id = ?", id).First(&user).Error
	if errDb != nil || user.ID == nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Data Found",
		"data":    user,
	})

}

func Store(ctx *gin.Context) {
	UserReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&UserReq); errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	user := new(models.Users)
	user.Username = &UserReq.Username
	user.Password = &UserReq.Password
	user.Role = &UserReq.Role
	user.MahasiswaID = &UserReq.MahasiswaID

	errDb := database.DB.Table("users").Create(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "can't create data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "data successfully created",
		"data":    user,
	})
}
