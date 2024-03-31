package auth_controllers

import (
	"Gin_Gorm_Api/database"
	models "Gin_Gorm_Api/models/UsersModels"
	"Gin_Gorm_Api/requests"
	"Gin_Gorm_Api/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {
	loginReq := new(requests.LoginRequest)
	if errReq := ctx.ShouldBind(&loginReq); errReq != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": errReq.Error(),
		})
		return
	}

	user := new(models.Users)
	errUser := database.DB.Table("users").Where("email = ?", loginReq.Email).Find(&user)

	if errUser != nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "Credentials not found or invalid",
		})
		return
	}

	//for compare password
	if loginReq.Password != "12345" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Credentials not found or invalid",
		})
		return
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error, can't generate token",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Login Success",
		"token":   token,
	})
}
