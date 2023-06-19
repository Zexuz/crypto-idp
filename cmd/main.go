package main

import (
	"context"
	"github.com/zexuz/crypto-idp/api"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type User struct {
	PublicAddress string `json:"publicAddress"`
	Nonce         int    `json:"nonce"`
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	errC := make(chan error)
	srv := api.StartServer(errC)

	// Listen for the interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped.")
}
