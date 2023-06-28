package middleware

import (
	"context"
	"errors"
	"github.com/zexuz/crypto-idp/api/render"
	"github.com/zexuz/crypto-idp/internal/jwt"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			render.Error(w, r, errors.New("authorization header is empty"), http.StatusUnauthorized)
			return
		}

		split := strings.Split(authHeader, " ")
		if len(split) != 2 {
			render.Error(w, r, errors.New("authorization header is not valid"), http.StatusUnauthorized)
			return
		}

		token, err := jwt.VerifyToken(split[1])
		if err != nil {
			render.Error(w, r, errors.New("token is not valid"), http.StatusUnauthorized)
			return
		}

		sub, err := token.Claims.GetSubject()
		if err != nil {
			render.Error(w, r, errors.New("could not get subject"), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "sub", sub)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
