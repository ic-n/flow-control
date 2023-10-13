package flow

import "github.com/pkg/errors"

// ErrNotSet represents an error indicating that a value is not set.
var ErrNotSet = errors.New("is not set")

// IsSet checks if the given value is set.
func IsSet[T comparable](v T) (T, error) {
	var i T

	if i == v {
		return v, ErrNotSet
	}

	return v, nil
}
