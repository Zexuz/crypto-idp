package errors

import "fmt"

type AddressMismatchError struct {
	RecoveredAddr string
	SignerAddr    string
}

func (e *AddressMismatchError) Error() string {
	return fmt.Sprintf("recovered address does not match signer address: got %s, want %s", e.RecoveredAddr, e.SignerAddr)
}
