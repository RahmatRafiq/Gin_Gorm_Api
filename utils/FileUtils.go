package FileUtils

import (
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func FileValidation(fileHeader *multipart.FileHeader, fileType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	log.Println("content-Type", contentType)
	result := false

	for _, typeFile := range fileType {
		if contentType == typeFile {
			result = true
			break
		}
	}
	return result
}

// func FileValidationByExtension(fileHeader *multipart.FileHeader, fileExtension []string) bool {
// 	extension := filepath.Ext(fileHeader.Filename)
// 	log.Println("extension", extension)
// 	result := false

// 	for _, typeFile := range fileExtension {
// 		if extension == "."+typeFile {
// 			result = true
// 			break
// 		}
// 	}
// 	return result
// }

func FileValidationByExtension(fileHeader *multipart.FileHeader, fileExtension []string) bool {
	extension := filepath.Ext(fileHeader.Filename)
	log.Println("extension", extension)

	for _, allowedExtension := range fileExtension {
		if extension == allowedExtension {
			return true
		}
	}

	return false
}

func RandomFileName(extensionFile string, prefix ...string) string {

	currentPrefix := "file"
	if len(prefix) > 0 && prefix[0] != "" {
		currentPrefix = prefix[0]
	}

	// currentPrefix := "file"
	// if len(prefix) > 0 {
	// 	if prefix[0] == "" {
	// 		currentPrefix = prefix[0]
	// 	}
	// } else {
	// 	prefix[0] = "file"
	// }

	currentTime := time.Now().UTC().Format("20060102T150405Z")
	filename := fmt.Sprintf("%s-%s-%s%s", currentPrefix, currentTime, RandomString(20), extensionFile)

	return filename

}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, filename string) bool {
	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", filename))

	if errUpload != nil {
		log.Println("Can't save file")
		return false
	} else {
		return true
	}

}
