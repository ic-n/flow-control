package flow

import "github.com/pkg/errors"

var ErrNotSet = errors.New("is not set")

func IsSet[T comparable](v T) (T, error) {
	var i T

	if i == v {
		return v, ErrNotSet
	}

	return v, nil
}
