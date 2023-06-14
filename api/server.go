package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zexuz/crypto-idp/api/nonce"
	"net/http"
)

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1/nonce", nonce.Routes())
	})

	println(fmt.Sprintf("Server started on port %d", 3000))
	http.ListenAndServe(":3000", r)
}
