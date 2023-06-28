package auth

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zexuz/crypto-idp/api/render"
	"github.com/zexuz/crypto-idp/internal/jwt"
	"net/http"
)

func (env *Env) PublicKey(writer http.ResponseWriter, request *http.Request) {

	key, err := jwt.GetPublicKey()
	if err != nil {
		logrus.WithError(err).Error("could not get public key")
		render.GenericError(writer, request)
		return
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		logrus.WithError(err).Error("could not marshal public key")
		render.GenericError(writer, request)
		return
	}

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	}
	bytes := pem.EncodeToMemory(block)

	render.OK(writer, request, struct {
		PublicKey string `json:"publicKey"`
	}{
		PublicKey: fmt.Sprintf("%s", bytes),
	})
}
