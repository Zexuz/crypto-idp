package auth

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zexuz/crypto-idp/internal/auth/errors"
	"github.com/zexuz/crypto-idp/internal/jwt"
)

const sigSize = 65

func ValidateAndCreateAccessToken(tokenStr string, signatureHex string) (string, error) {
	receivedToken, err := jwt.VerifyToken(tokenStr)
	if err != nil {
		return "", fmt.Errorf("could not verify token: %w", err)
	}

	signerAddress, err := jwt.GetSubClaimsFromToken(receivedToken)
	if signerAddress == "" {
		return "", fmt.Errorf("could not get signer address from token: %w", err)
	}

	if err := validateSigner(signatureHex, tokenStr, signerAddress); err != nil {
		return "", err
	}

	token, err := jwt.GetNewToken(signerAddress)
	if err != nil {
		return "", fmt.Errorf("could not create new token: %w", err)
	}

	return token, nil
}

func validateSigner(signatureHex string, tokenStr string, signerAddress string) error {
	signature, err := hex.DecodeString(signatureHex[2:]) // Convert from hexadecimal string to byte slice, omitting the "0x" prefix
	if err != nil {
		return &errors.SignatureDecodeError{Signature: signatureHex, Err: err}
	}

	if len(signature) != sigSize {
		return &errors.SignatureSizeError{ExpectedSize: sigSize, ActualSize: len(signature)}
	}
	// Add prefix to original message
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(tokenStr), tokenStr)

	// Hash the prefixed message
	hash := crypto.Keccak256Hash([]byte(prefixedMessage))

	// Convert to EIP 155 format for signature
	signature[64] -= 27
	// Recover the public key from the signature
	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return &errors.PublicKeyRecoveryError{Err: err}
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// Verify the recovered address with the signer's address
	if recoveredAddr.Hex() != signerAddress {
		return &errors.AddressMismatchError{RecoveredAddr: recoveredAddr.Hex(), SignerAddr: signerAddress}
	}

	return nil
}
