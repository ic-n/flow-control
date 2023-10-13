package flow

// Default creates a Functor that, when applied to a value and an error, returns the original
// value if there is no error; otherwise, it returns the provided default value along with the error.
// It allows for setting a default value in case of an error.
func Default[T any](d T) Functor[T] {
	return func(o T, err error) (T, error) {
		if err != nil {
			return d, err
		}

		return o, err
	}
}

// DefaultExcuse creates a Functor that, when applied to a value and an error, returns the original
// value if there is no error; otherwise, it returns the provided default value with the error omitted.
// It allows for setting a default value in case of an error without carrying the error forward.
func DefaultExcuse[T any](d T) Functor[T] {
	return func(o T, err error) (T, error) {
		if err != nil {
			return d, nil
		}

		return o, nil
	}
}
