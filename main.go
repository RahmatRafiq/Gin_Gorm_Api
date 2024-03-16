package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Tambahkan rute-rute API di sini

	r.Run(":8080")
}
