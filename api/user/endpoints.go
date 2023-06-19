package user

import (
	"github.com/go-chi/chi/v5"
)

type Env struct {
}

func Routes() *chi.Mux {
	env := &Env{}
	r := chi.NewRouter()

	r.Post("/", env.Me)

	return r
}
