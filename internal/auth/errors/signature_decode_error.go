package errors

import "fmt"

type SignatureDecodeError struct {
	Signature string
	Err       error
}

func (e *SignatureDecodeError) Error() string {
	return fmt.Sprintf("failed to decode signature %s: %v", e.Signature, e.Err)
}

func (e *SignatureDecodeError) Unwrap() error {
	return e.Err
}
