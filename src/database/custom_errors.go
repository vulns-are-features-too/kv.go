package database

import (
	"fmt"
)

type NotFoundError struct {
	key string
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("key not found: %s", err.key)
}

func notFoundError(key string) error {
	return NotFoundError{key: key}
}

type IncompatibleTypeError struct {
	expected string
	actual   string
}

func (err IncompatibleTypeError) Error() string {
	return fmt.Sprintf("Incompatible types: expected %s, got %s", err.expected, err.actual)
}

func incompatibleTypeError(expected string, actual string) error {
	return IncompatibleTypeError{expected: expected, actual: actual}
}
