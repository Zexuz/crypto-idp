package user

import (
	"github.com/zexuz/crypto-idp/api/render"
	"net/http"
)

func (env *Env) Me(writer http.ResponseWriter, request *http.Request) {
	sub := request.Context().Value("sub").(string)

	render.OK(writer, request, sub)
}
