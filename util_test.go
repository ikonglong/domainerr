package domainerr

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCheckArgument(t *testing.T) {
	numRange := []int{1, 10}
	n := 11
	err := CheckArgument(n >= numRange[0] && n <= numRange[1], "range %v not include %d", numRange, n)
	assert.Equal(t, "illegal argument: range [1 10] not include 11", err.Error())
}

func TestChainMsg(t *testing.T) {
	assert.Equal(t, "", ChainMsg(nil))

	e1 := NewInternalError().WithMessage("error a").Build()
	assert.Equal(t, "error a", ChainMsg(e1))

	e2 := NewInternalError().WithMessage("error a").
		WithCause(fmt.Errorf("error b")).
		Build()
	assert.Equal(t, "error a -> error b", ChainMsg(e2))

	e3 := NewInternalError().WithMessage("error a").
		WithCause(NewUnknownError().WithMessage("error b").WithCause(fmt.Errorf("error c")).Build()).
		Build()
	assert.Equal(t, "error a -> error b -> error c", ChainMsg(e3))

	var nilDomainErr *Error
	e4 := NewInternalError().WithMessage("error a").
		WithCause(nilDomainErr).Build()
	assert.Equal(t, "error a", ChainMsg(e4))

	e5 := NewInternalError().WithMessage("error a").
		WithCause(nil).Build()
	assert.Equal(t, "error a", ChainMsg(e5))

	var nilErr error = nil
	e6 := NewInternalError().WithMessage("error a").
		WithCause(nilErr).Build()
	assert.Equal(t, "error a", ChainMsg(e6))

	e7 := NewInternalError().WithMessage("error a").
		WithCause(errors.WithStack(fmt.Errorf("error b"))).Build()
	assert.Equal(t, "error a -> error b -> error b", ChainMsg(e7))

	e8 := NewInternalError().WithMessage("error a").
		WithCause(fmt.Errorf("wrap error: %w", fmt.Errorf("error c"))).Build()
	assert.Equal(t, "error a -> wrap error: error c", ChainMsg(e8))
}
