package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"path"
	"runtime"
)

func currentFileDir() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("could not get current file directory")
	}
	return path.Dir(filename), nil
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	dir, err := currentFileDir()
	if err != nil {
		return nil, err
	}

	assetFilePath := path.Join(dir, "../../assets/keys/private_key.pem")

	privateKeyBytes, err := os.ReadFile(assetFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	return privateKey, nil
}

func getPublicKey() (*rsa.PublicKey, error) {
	dir, err := currentFileDir()
	if err != nil {
		return nil, err
	}

	assetFilePath := path.Join(dir, "../../assets/keys/public_key.pem")

	publicKeyBytes, err := os.ReadFile(assetFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read public key: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	return pubKey, nil
}
