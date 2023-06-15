package main

import (
	"github.com/zexuz/crypto-idp/api"
	"log"
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

	api.StartServer()
}
