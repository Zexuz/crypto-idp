package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/zexuz/crypto-idp/internal/database"
)

type Env struct {
	db *database.UserDatabaseService
}

func Routes(db *database.UserDatabaseService) *chi.Mux {
	env := &Env{
		db: db,
	}
	r := chi.NewRouter()

	r.Post("/", env.Me)

	return r
}
