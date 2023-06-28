package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func createTokenWithClaims(claims jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	privKey, err := getPrivateKey()
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(privKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
