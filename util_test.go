package domainerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckArgument(t *testing.T) {
	numRange := []int{1, 10}
	n := 11
	err := CheckArgument(n >= numRange[0] && n <= numRange[1], "range %v not include %d", numRange, n)
	assert.Equal(t, "illegal argument: range [1 10] not include 11", err.Error())
}
