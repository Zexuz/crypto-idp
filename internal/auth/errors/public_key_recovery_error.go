package errors

import "fmt"

type PublicKeyRecoveryError struct {
	Err error
}

func (e *PublicKeyRecoveryError) Error() string {
	return fmt.Sprintf("failed to recover public key: %v", e.Err)
}

func (e *PublicKeyRecoveryError) Unwrap() error {
	return e.Err
}
