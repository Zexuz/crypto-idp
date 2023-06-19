package jwt

import "github.com/golang-jwt/jwt/v5"

func GetSubClaimsFromToken(token *jwt.Token) (string, error) {

	nonce, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	return nonce, nil
}
