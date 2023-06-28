package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/zexuz/crypto-idp/api/auth"
	"github.com/zexuz/crypto-idp/api/user"
	"net/http"
	"regexp"
)

func StartServer(errC chan error) *http.Server {
	r := chi.NewRouter()

	addMiddleware(r)
	addRoutes(r)

	server := &http.Server{Addr: ":3000", Handler: r}

	_, cancel := context.WithCancel(context.Background())
	go func() {
		println(fmt.Sprintf("Server started on port %d", 3000))
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			cancel()
			errC <- errors.New("ServerHTTP: server.http.ListenAndServe(): " + err.Error())
		}
	}()

	return server
}

func addRoutes(r *chi.Mux) {

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1/auth", auth.Routes())
		r.Mount("/v1/me", user.Routes())
	})
}

func addMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

func AllowOriginFunc(_ *http.Request, origin string) bool {
	if matched, _ := regexp.MatchString("http://localhost:[0-9]+", origin); matched {
		return true
	}

	return false
}
