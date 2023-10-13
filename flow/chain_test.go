package flow_test

import (
	"testing"

	"github.com/ic-n/flow-control/flow"
	"github.com/stretchr/testify/assert"
)

type Dummy struct {
	ID int
}

func TestEval(t *testing.T) {
	c := flow.New[*Dummy]()
	c.Maybe(flow.IsSet[*Dummy])
	c.Pipe(flow.DefaultExcuse(&Dummy{ID: -1}))

	result, err := c.Eval(&Dummy{ID: 1})
	assert.NoError(t, err)
	assert.Equal(t, 1, result.ID)

	result, err = c.Eval(nil)
	assert.NoError(t, err)
	assert.Equal(t, -1, result.ID)
}
