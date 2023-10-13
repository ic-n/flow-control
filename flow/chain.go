package flow

import "github.com/pkg/errors"

type Functor[T any] func(T, error) (T, error)

type Chainlink[T any] struct {
	Functor  Functor[T]
	Previous *Chainlink[T]
}

func New[T any]() *Chainlink[T] {
	return &Chainlink[T]{
		Functor: func(t T, err error) (T, error) { return t, err },
	}
}

func (c *Chainlink[T]) Pipe(fn Functor[T]) {
	*c = Chainlink[T]{
		Functor:  fn,
		Previous: &Chainlink[T]{c.Functor, c.Previous},
	}
}

func (c *Chainlink[T]) Maybe(fn func(T) (T, error)) {
	*c = Chainlink[T]{
		Functor: func(v T, err error) (T, error) {
			if err != nil {
				return v, err
			}

			return fn(v)
		},
		Previous: &Chainlink[T]{c.Functor, c.Previous},
	}
}

func (c *Chainlink[T]) Finally(fn func(T) (T, error)) {
	*c = Chainlink[T]{
		Functor: func(v T, err error) (T, error) {
			v2, err2 := fn(v)
			if err2 != nil {
				return v2, errors.Wrapf(err2, "with original error %s", err.Error())
			}

			return v2, err
		},
		Previous: &Chainlink[T]{c.Functor, c.Previous},
	}
}

func (c *Chainlink[T]) Eval(v T) (T, error) {
	var (
		r     = v
		err   error
		chain = c.ops()
	)

	for i := len(chain) - 1; i >= 0; i-- {
		r, err = chain[i](r, err)
	}

	return r, err
}

func (c *Chainlink[T]) ops() []Functor[T] {
	chain := make([]Functor[T], 0)
	current := *c

	for {
		chain = append(chain, current.Functor)
		if current.Previous == nil {
			break
		}

		current = *current.Previous
	}

	return chain
}
