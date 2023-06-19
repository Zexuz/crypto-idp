package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GetNewNonceToken(publicAddress string, nonce string) (string, error) {
	exp := time.Now().Add(time.Minute * 5)

	claims := jwt.MapClaims{
		"sub":   publicAddress,
		"exp":   exp.Unix(),
		"nonce": nonce,
	}

	return createTokenWithClaims(claims)
}
