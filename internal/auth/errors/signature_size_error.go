package errors

import "fmt"

type SignatureSizeError struct {
	ExpectedSize int
	ActualSize   int
}

func (e *SignatureSizeError) Error() string {
	return fmt.Sprintf("wrong size for signature: got %d, want %d", e.ActualSize, e.ExpectedSize)
}
