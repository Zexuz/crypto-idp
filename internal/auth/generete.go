package auth

import (
	"fmt"
	"github.com/zexuz/crypto-idp/internal/jwt"
	"github.com/zexuz/crypto-idp/internal/nonce"
)

func GenerateTokenToSign(publicAddress string) (string, error) {
	generator := nonce.NewRandomNonceGeneratorDefault()
	n, err := generator.GenerateNonce()
	if err != nil {
		return "", fmt.Errorf("could not generate nonce: %w", err)
	}

	token, err := jwt.GetNewNonceToken(publicAddress, n)
	if err != nil {
		return "", fmt.Errorf("could not generate token: %w", err)
	}
	return token, nil
}
