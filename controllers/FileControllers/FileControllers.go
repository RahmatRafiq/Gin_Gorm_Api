package FileControllers

import (
	"Gin_Gorm_Api/costanta"
	utils "Gin_Gorm_Api/utils"
	"net/http"
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
	fileExtension := []string{".png", ".jpg", ".jpeg"}
	isFileValidated := utils.FileValidationByExtension(fileHeader, fileExtension)

	if !isFileValidated {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File type not allowed",
		})
		return
	}

	extensionFile := filepath.Ext(fileHeader.Filename)
	filename := utils.RandomFileName(extensionFile)

	isSaved := utils.SaveFile(ctx, fileHeader, filename)

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

func HandleDestroyFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	if filename == "" {
		ctx.JSON(400, gin.H{
			"message": "Filename is required",
		})
	}

	err := utils.DestroyFile(costanta.DIR_FILE + filename)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "File Deleted",
	})
}
