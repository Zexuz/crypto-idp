package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type UnexpectedSigningMethodError struct {
	SigningMethod string
}

func (e *UnexpectedSigningMethodError) Error() string {
	return fmt.Sprintf("unexpected signing method: %v", e.SigningMethod)
}

type TokenVerificationError struct {
	Err error
}

func (e *TokenVerificationError) Error() string {
	return fmt.Sprintf("failed to verify token: %v", e.Err)
}

func (e *TokenVerificationError) Unwrap() error {
	return e.Err
}

type PublicKeyRecoveryError struct {
	Err error
}

func (e *PublicKeyRecoveryError) Error() string {
	return fmt.Sprintf("failed to recover public key: %v", e.Err)
}

func (e *PublicKeyRecoveryError) Unwrap() error {
	return e.Err
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, &UnexpectedSigningMethodError{token.Header["alg"].(string)}
		}

		pubKey, err := GetPublicKey()
		if err != nil {
			return "", &PublicKeyRecoveryError{err}
		}

		return pubKey, nil
	})

	if err != nil {
		return nil, &TokenVerificationError{err}
	}

	if !token.Valid {
		return nil, &TokenVerificationError{errors.New("token is not valid")}
	}

	return token, nil
}
