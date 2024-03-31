package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(gin *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, gin)

}
