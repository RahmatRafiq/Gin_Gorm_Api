package FileControllers

import (
	"fmt"

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

	// file, errFile := fileHeader.Open()
	// 	if errFile != nil {
	// 		panic(errFile)
	// 	}

	// extensionFile := filepath.Ext(fileHeader.Filename)
	// filename :=

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", fileHeader.Filename))

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
