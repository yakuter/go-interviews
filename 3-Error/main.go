package main

import (
	"errors"
	"fmt"
)

// CustomError is a custom error type that implements the error interface and Unwrap method
type CustomError struct {
	Code    int
	Message string
	Err     error // Embed the underlying error for unwrapping
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// String method for CustomError to satisfy the Stringer interface
func (e CustomError) String() string {
	return e.Error()
}

// Unwrap method for CustomError to satisfy the Unwrap interface
func (e CustomError) Unwrap() error {
	return e.Err
}

// Function that returns a custom error or wraps an existing error with additional context
func process(data int) error {
	if data < 0 {
		return &CustomError{Code: 400, Message: "Invalid input data", Err: errors.New("input error")}
	}
	return fmt.Errorf("processing failed for data: %d", data)
}

func main() {
	// Example of custom error type
	err := process(-1)
	if err != nil {
		fmt.Println("Custom Error:", err)
	}

	// Example of error wrapping with fmt.Errorf
	wrappedErr := fmt.Errorf("additional context: %w", err)
	fmt.Println("Wrapped Error:", wrappedErr)

	// Example of errors.Is
	if errors.Is(wrappedErr, &CustomError{}) {
		fmt.Println("Wrapped error is of type CustomError")
	}

	// Example of errors.As
	var ce *CustomError
	if errors.As(wrappedErr, &ce) {
		fmt.Printf("Custom Error Code: %d, Message: %s\n", ce.Code, ce.Message)
	}

	// Example of error unwrapping using Unwrap method
	if unwrappedErr := errors.Unwrap(wrappedErr); unwrappedErr != nil {
		fmt.Println("Unwrapped Error:", unwrappedErr)
	}
}
