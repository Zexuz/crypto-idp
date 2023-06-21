package user

import (
	"github.com/zexuz/crypto-idp/api/types"
	"github.com/zexuz/crypto-idp/internal/jwt"
	"net/http"
	"strings"
)

func (env *Env) Me(writer http.ResponseWriter, request *http.Request) {

	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		types.FailureResponse("Authorization header is empty", writer, request, http.StatusUnauthorized)
		return
	}

	split := strings.Split(authHeader, " ")
	if len(split) != 2 {
		types.FailureResponse("Authorization header is not valid", writer, request, http.StatusUnauthorized)
		return
	}

	token, err := jwt.VerifyToken(split[1])
	if err != nil {
		types.FailureResponse("Token is not valid", writer, request, http.StatusUnauthorized)
		return
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		types.FailureResponse("Could not get subject", writer, request, http.StatusInternalServerError)
		return
	}

	types.SuccessResponse(sub, writer, request)
}
