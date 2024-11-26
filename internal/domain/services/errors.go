package services

import "fmt"

func ErrInvalidInput(message string) error {
	return fmt.Errorf("invalid input: %s", message)
}

func ErrNotFound(message string) error {
	return fmt.Errorf("not found: %s", message)
}
