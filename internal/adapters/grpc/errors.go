package grpc

import "fmt"

func ErrTaskNotFound(message string) error {
	return fmt.Errorf("not found: %s", message)
}
