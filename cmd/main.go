package main

import (
	"github.com/zexuz/crypto-idp/api"
)

type User struct {
	PublicAddress string `json:"publicAddress"`
	Nonce         int    `json:"nonce"`
}

func main() {
	api.StartServer()
}
