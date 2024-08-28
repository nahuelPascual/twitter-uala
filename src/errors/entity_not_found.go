package errors

import "fmt"

type EntityNotFoundError struct {
	Entity string
	Key    string
	Value  any
}

func (err EntityNotFoundError) Error() string {
	return fmt.Sprintf("not found %s with %s=%v", err.Entity, err.Key, err.Value)
}
