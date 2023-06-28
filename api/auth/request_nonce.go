package auth

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zexuz/crypto-idp/api/render"
	"github.com/zexuz/crypto-idp/internal/auth"
	"net/http"
)

type RequestNonceResponse struct {
	Token string `json:"token"`
}

const publicAddressQueryParam = "publicAddress"

func (env *Env) RequestNonce(writer http.ResponseWriter, request *http.Request) {
	publicAddress := request.URL.Query().Get(publicAddressQueryParam)
	if publicAddress == "" {
		err := fmt.Errorf("missing queryString %s", publicAddressQueryParam)
		render.Error(writer, request, err, http.StatusBadRequest)
		return
	}

	token, err := auth.GenerateTokenToSign(publicAddress)
	if err != nil {
		logrus.WithError(err).Error("Could not generate token")
		render.GenericError(writer, request)
		return
	}

	response := RequestNonceResponse{
		Token: token,
	}

	render.OK(writer, request, response)
}
