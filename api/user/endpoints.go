package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/zexuz/crypto-idp/api/middleware"
)

type Env struct {
}

func Routes() *chi.Mux {
	env := &Env{}
	r := chi.NewRouter()

	r.Use(middleware.Auth)

	r.Post("/", env.Me)

	return r
}
