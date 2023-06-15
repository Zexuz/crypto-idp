package user

import (
	"github.com/zexuz/crypto-idp/api/types"
	"github.com/zexuz/crypto-idp/internal/jwt"
	"net/http"
	"strings"
)

func (env *Env) Me(writer http.ResponseWriter, request *http.Request) {

	authHeader := request.Header.Get("Authorization")

	split := strings.Split(authHeader, " ")[1]

	token, err := jwt.VerifyToken(split)
	if err != nil {
		types.FailureResponse("Token is not valid", writer, request)
		return
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		types.FailureResponse("Could not get subject", writer, request)
		return
	}

	types.SuccessResponse(sub, writer, request)
}
