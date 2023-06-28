package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GetNewToken(publicAddress string) (string, error) {
	exp := time.Now().Add(time.Hour)

	claims := jwt.MapClaims{
		"sub": publicAddress,
		"exp": exp.Unix(),
	}

	return createTokenWithClaims(claims)
}
