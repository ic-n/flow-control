// Package flow defines a simple chaining mechanism for functions that operate on values
// and errors. It allows you to create chains of functions (Functors) and apply them to a value.
package flow

import "errors"

// Functor represents a elementary individual function, part of conversion chain.
type Functor[T any] func(T, error) (T, error)

// Chainlink is a structure that represents a link in the chain of functions (Functors).
type Chainlink[T any] struct {
	// Functor is function to be applied.
	Functor Functor[T]
	// Previous is previous link in the functor chain.
	Previous *Chainlink[T]
}

// New creates a new Chainlink with initial no-op functor.
func New[T any]() *Chainlink[T] {
	return &Chainlink[T]{
		Functor: func(t T, err error) (T, error) { return t, err },
	}
}

// Pipe appends a new Functor to the chain.
func (c *Chainlink[T]) Pipe(fn Functor[T]) {
	*c = Chainlink[T]{
		Functor:  fn,
		Previous: &Chainlink[T]{c.Functor, c.Previous},
	}
}

// Maybe appends a new Functor to the chain that conditionally applies a function if there is no error.
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

// Finally appends a new Functor to the chain that applies a function regardless of the error status,
// and wraps any new error with the original error.
func (c *Chainlink[T]) Finally(fn func(T) (T, error)) {
	*c = Chainlink[T]{
		Functor: func(v T, err error) (T, error) {
			v2, err2 := fn(v)
			if err2 != nil {
				return v2, errors.Join(err, err2)
			}

			return v2, err
		},
		Previous: &Chainlink[T]{c.Functor, c.Previous},
	}
}

// Eval applies the chain of Functors to the input value and returns the final result and error.
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
