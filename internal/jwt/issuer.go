package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

type Token struct {
}

const (
	publicAddressClaim = "publicAddressClaim"
)

// TODO Read thi from the env
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

func GetNewToken(publicAddress string) (string, error) {

	exp := time.Now().Add(time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": publicAddress,
		"exp": exp.Unix(),
	})

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

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		pubKey, err := getPublicKey()
		if err != nil {
			return "", err
		}

		return pubKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return token, nil
}
