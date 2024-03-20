package UserController

import (
	"Gin_Gorm_Api/database"
	models "Gin_Gorm_Api/models/UsersModels"
	"Gin_Gorm_Api/requests"
	"Gin_Gorm_Api/responses"
	"log"
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
	userReq := new(requests.UserRequest)

	if err := ctx.ShouldBind(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	userEmailExist := new(models.Users)
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)
	if userEmailExist.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exists",
		})
		return
	}

	user := models.Users{
		Username:    &userReq.Username,
		Email:       &userReq.Email,
		Password:    &userReq.Password,
		Role:        &userReq.Role,
		MahasiswaID: &userReq.MahasiswaID,
	}

	if err := database.DB.Table("users").Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"data":    user,
	})
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.Users)
	userReq := new(requests.UserRequest)
	userEmailExist := new(models.Users)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDB := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDB != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "User data not found",
		})
		return
	}
	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error
	if errUserEmailExist != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(400, gin.H{
			"message": "Email already used",
		})
		return
	}

	user.Username = &userReq.Username
	user.Email = &userReq.Email
	user.Password = &userReq.Password
	user.Role = &userReq.Role
	user.MahasiswaID = &userReq.MahasiswaID

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "can't update data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Data updated successfully",
		"data":    user,
	})
}

func Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.Users)

	errFind := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errFind != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error":   errFind.Error(),
		})
		return
	}
	log.Println("user", user)

	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "User data not found",
		})
		return
	}

	errDB := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&models.Users{}).Error
	if errDB != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error":   errDB.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Data deleted successfully",
	})
}
