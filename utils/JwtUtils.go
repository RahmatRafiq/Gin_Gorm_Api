package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = "SECRET_KEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webtoken, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}
	return webtoken, nil
}

func VerifiyToken(tokenString string) (*jwt.Token, error) {
	tokenJwt, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)
		if isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret_key), nil
	})
	if err != nil {
		return nil, err
	}

	return tokenJwt, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifiyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOK := token.Claims.(jwt.MapClaims)
	if isOK && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Invalid Token")

	}
}
