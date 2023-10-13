package flow

func Default[T any](d T) Functor[T] {
	return func(o T, err error) (T, error) {
		if err != nil {
			return d, err
		}

		return o, err
	}
}

func DefaultExcuse[T any](d T) Functor[T] {
	return func(o T, err error) (T, error) {
		if err != nil {
			return d, nil
		}

		return o, nil
	}
}
