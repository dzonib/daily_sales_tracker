package util

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	payload := jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Hour * 24).Unix())),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
