package auth

import (
	"errors"
	chirender "github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"github.com/zexuz/crypto-idp/api/render"
	"github.com/zexuz/crypto-idp/internal/auth"
	cerr "github.com/zexuz/crypto-idp/internal/auth/errors"
	"net/http"
)

type CallbackResponse struct {
	Jwt string `json:"jwt"`
}

type CallbackRequest struct {
	Signature string `json:"signature"`
	JwtToken  string `json:"jwtToken"`
}

func (d *CallbackRequest) Bind(_ *http.Request) error {
	return nil
}

func (env *Env) Callback(writer http.ResponseWriter, request *http.Request) {
	requestBody := &CallbackRequest{}
	if err := chirender.DecodeJSON(request.Body, requestBody); err != nil {
		render.Error(writer, request, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}

	token, err := auth.ValidateAndCreateAccessToken(requestBody.JwtToken, requestBody.Signature)
	if err != nil {
		logrus.WithError(err).Error("Could not handle callback")
		// TODO use errors.Is, and create a base error type
		switch err.(type) {
		case *cerr.SignatureDecodeError:
		case *cerr.SignatureSizeError:
		case *cerr.PublicKeyRecoveryError:
		case *cerr.AddressMismatchError:
			render.Error(writer, request, errors.New("error validate signature"), http.StatusUnauthorized)
		default:
			render.GenericError(writer, request)
		}

		return
	}

	response := CallbackResponse{
		Jwt: token,
	}

	render.OK(writer, request, response)
}
