package nonce

import (
	"github.com/go-chi/chi/v5"
	"github.com/zexuz/crypto-idp/internal/database"
)

type Env struct {
	db *database.UserDatabaseService
}

type Response struct {
	Nonce string `json:"nonce"`
}

func Routes(db *database.UserDatabaseService) *chi.Mux {
	env := &Env{
		db: db,
	}
	r := chi.NewRouter()

	r.Get("/", env.RequestNonce)
	r.Post("/", env.Callback)

	return r
}
