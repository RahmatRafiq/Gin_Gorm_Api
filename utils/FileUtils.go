package FileUtils

import (
	"math/rand"
	"mime/multipart"
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
	result := false

	for _, typeFile := range fileType {
		if contentType == typeFile {
			result = true
			break
		}
	}
	return result
}
