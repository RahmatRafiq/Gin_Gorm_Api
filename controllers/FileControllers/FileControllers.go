package FileControllers

import "github.com/gin-gonic/gin"

func UploadFile(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "File Uploaded",
	})
}
