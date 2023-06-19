package nonce

import (
	"github.com/go-chi/chi/v5"
)

type Env struct {
}

type Response struct {
	Nonce string `json:"nonce"`
}

func Routes() *chi.Mux {
	env := &Env{}
	r := chi.NewRouter()

	r.Get("/", env.RequestNonce)
	r.Post("/", env.Callback)

	return r
}
