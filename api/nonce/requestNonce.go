package nonce

import (
	"fmt"
	"github.com/zexuz/crypto-idp/api/types"
	"github.com/zexuz/crypto-idp/internal/nonce"
	"net/http"
)

type RequestNonceResponse struct {
	Nonce string `json:"nonce"`
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

	if err = env.db.SetUserNonce(publicAddress, n); err != nil {
		types.FailureResponse("Could not set nonce", writer, request)
		return
	}

	response := RequestNonceResponse{
		Nonce: n,
	}

	types.SuccessResponse(response, writer, request)
}
