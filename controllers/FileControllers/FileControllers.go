package FileControllers

import (
	FileUtils "Gin_Gorm_Api/utils"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {

	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "File is required",
		})
		return
	}

	extensionFile := filepath.Ext(fileHeader.Filename)
	currentTime := time.Now().UTC().Format("20060102T150405Z")
	filename := fmt.Sprintf("%s-%s%s", currentTime, FileUtils.RandomString(20), extensionFile)
	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", filename))

	if errUpload != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error, Failed to upload file",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "File Uploaded",
	})
}
