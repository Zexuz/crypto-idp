package nonce

import (
	"crypto/rand"
	"math/big"
)

type GenerationSettings struct {
	Length int
	Chars  string
}

type Generator interface {
	GenerateNonce() string
}

type RandomNonceGenerator struct {
	settings GenerationSettings
}

func NewRandomNonceGeneratorDefault() *RandomNonceGenerator {
	// TODO get config from env vars
	settings := GenerationSettings{
		Length: 32,
		Chars:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
	return NewRandomNonceGenerator(settings)
}

func NewRandomNonceGenerator(settings GenerationSettings) *RandomNonceGenerator {
	return &RandomNonceGenerator{settings: settings}
}

func (r *RandomNonceGenerator) GenerateNonce() (string, error) {
	nonce := make([]byte, r.settings.Length)
	for i := range nonce {
		max := big.NewInt(int64(len(r.settings.Chars)))
		number, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		nonce[i] = r.settings.Chars[number.Int64()]
	}
	return string(nonce), nil
}
