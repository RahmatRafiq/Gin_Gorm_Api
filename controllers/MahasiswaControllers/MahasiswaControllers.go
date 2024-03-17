package MahasiswaControllers

import "github.com/gin-gonic/gin"

func GetAllMahasiswa(ctx *gin.Context) {

	isValidated := true

	if !isValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "Bad Request Something Field not Valid",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Hello Mahasiswa",
	})

}
