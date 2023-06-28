package auth

import (
	"github.com/go-chi/chi/v5"
)

type Env struct {
}

type Response struct {
	Nonce string `json:"auth"`
}

func Routes() *chi.Mux {
	env := &Env{}
	r := chi.NewRouter()

	r.Get("/", env.RequestNonce)
	r.Post("/", env.Callback)
	r.Get("/public-key", env.PublicKey)

	return r
}
