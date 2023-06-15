package nonce

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-chi/render"
	"github.com/zexuz/crypto-idp/api/types"
	"github.com/zexuz/crypto-idp/internal/jwt"
	"log"
	"net/http"
)

type CallbackResponse struct {
	Jwt string `json:"jwt"`
}

type CallbackRequest struct {
	Signature     string `json:"signature"`
	PublicAddress string `json:"publicAddress"`
}

func (d *CallbackRequest) Bind(request *http.Request) error {
	return nil
}

// message is "publicAddress:nonce"
func (env *Env) Callback(writer http.ResponseWriter, request *http.Request) {
	// TODO get the publicAddress and nonce from the request body

	requestBody := &CallbackRequest{}
	if err := render.DecodeJSON(request.Body, requestBody); err != nil {
		types.FailureResponse("Could not decode request body", writer, request)
		return
	}

	nonce, err := env.db.GetUserNonce(requestBody.PublicAddress)
	if err != nil {
		types.FailureResponse("Could not get nonce", writer, request)
		return
	}

	if nonce == "" {
		types.FailureResponse("Nonce is empty", writer, request)
		return
	}

	originalMessage := nonce

	// This is the signer's address
	signerAddress := requestBody.PublicAddress

	// This is your signed message (retrieved from MetaMask)
	signatureHex := requestBody.Signature                // This is a hexadecimal string
	signature, err := hex.DecodeString(signatureHex[2:]) // Convert from hexadecimal string to byte slice, omitting the "0x" prefix
	if err != nil {
		msg := fmt.Sprintf("Invalid signature format: %v", err)
		types.FailureResponse(msg, writer, request)
		return
	}

	// Add prefix to original message
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(originalMessage), originalMessage)

	// Hash the prefixed message
	hash := crypto.Keccak256Hash([]byte(prefixedMessage))

	if len(signature) != 65 {
		msg := fmt.Sprintf("wrong size for signature: got %d, want 65", len(signature))
		types.FailureResponse(msg, writer, request)
		return
	}

	// Convert to EIP 155 format for signature
	signature[64] -= 27

	// Recover the public key from the signature
	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatalf("failed to recover public key: %v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// Verify the recovered address with the signer's address
	if recoveredAddr.Hex() != signerAddress {
		msg := fmt.Sprintf("recovered address does not match signer address: got %s, want %s", recoveredAddr.Hex(), signerAddress)
		types.FailureResponse(msg, writer, request)
		return
	}

	// TODO create a jwt token and return it

	token, err := jwt.GetNewToken(requestBody.PublicAddress)
	if err != nil {
		types.FailureResponse("Could not create token", writer, request)
		return
	}

	response := CallbackResponse{
		Jwt: token,
	}

	types.SuccessResponse(response, writer, request)
}
