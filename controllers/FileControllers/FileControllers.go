package FileControllers

import (
	FileUtils "Gin_Gorm_Api/utils"
	"path/filepath"

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

	fileExtension := []string{"png"}
	isFileValidated := FileUtils.FileValidation(fileHeader, fileExtension)

	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file not allowed ",
		})
	}

	// fileType := []string{"image/img"}
	// isFileValidated := FileUtils.FileValidation(fileHeader, fileType)

	// if !isFileValidated {
	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "file not allowed ",
	// 	})
	// }

	extensionFile := filepath.Ext(fileHeader.Filename)
	filename := FileUtils.RandomFileName(extensionFile)
	// filename := FileUtils.RandomFileName(extensionFile, "file")

	isSaved := FileUtils.SaveFile(ctx, fileHeader, filename)

	if !isSaved {
		ctx.JSON(500, gin.H{
			"message": "internal server error, can't save file",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "File Uploaded",
	})
}
