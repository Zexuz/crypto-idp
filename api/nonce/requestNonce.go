package nonce

import (
	"fmt"
	"github.com/zexuz/crypto-idp/api/types"
	"github.com/zexuz/crypto-idp/internal/jwt"
	"github.com/zexuz/crypto-idp/internal/nonce"
	"net/http"
)

type RequestNonceResponse struct {
	Token string `json:"token"`
}

const publicAddressQueryParam = "publicAddress"

func (env *Env) RequestNonce(writer http.ResponseWriter, request *http.Request) {
	publicAddress := request.URL.Query().Get(publicAddressQueryParam)
	if publicAddress == "" {
		types.FailureResponse(fmt.Sprintf("Missing queryString %s", publicAddressQueryParam), writer, request)
		return
	}

	generator := nonce.NewRandomNonceGeneratorDefault()
	n, err := generator.GenerateNonce()
	if err != nil {
		println(err.Error())
		types.FailureResponse("Could not generate nonce", writer, request)
		return
	}

	token, err := jwt.GetNewNonceToken(publicAddress, n)
	if err != nil {
		types.FailureResponse("Could not generate token", writer, request)
		return
	}

	response := RequestNonceResponse{
		Token: token,
	}

	types.SuccessResponse(response, writer, request)
}
