package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zexuz/crypto-idp/api/nonce"
	"github.com/zexuz/crypto-idp/api/user"
	"github.com/zexuz/crypto-idp/internal/database"
	"net/http"
)

func StartServer(errC chan error) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db := database.NewDatabase()

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1/nonce", nonce.Routes(db))
		r.Mount("/v1/me", user.Routes(db))
	})

	server := &http.Server{Addr: ":3000", Handler: r}

	_, cancel := context.WithCancel(context.Background())
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			cancel()
			errC <- errors.New("ServerHTTP: server.http.ListenAndServe(): " + err.Error())
		}
		println(fmt.Sprintf("Server started on port %d", 3000))
	}()

	return server

}
