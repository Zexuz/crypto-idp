package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"path"
	"runtime"
)

func currentFileDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Could not get current file info")
	}
	dir := path.Dir(filename)
	return dir
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	assetFilePath := path.Join(currentFileDir(), "../../assets/keys/private_key.pem")

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
	assetFilePath := path.Join(currentFileDir(), "../../assets/keys/public_key.pem")

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
