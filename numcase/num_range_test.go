package numcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNumRange_IllegalStart(t *testing.T) {
	_, err := NewNumRange(-1, 10)
	assert.EqualError(t, err, "illegal argument: start < 0")
}

func TestNewNumRange_IllegalEnd(t *testing.T) {
	_, err := NewNumRange(1, -10)
	assert.EqualError(t, err, "illegal argument: end < 0")
}

func TestNewNumRange_StartGtEnd(t *testing.T) {
	_, err := NewNumRange(10, 1)
	assert.EqualError(t, err, "illegal argument: end < start")
}

func TestNewNumRange_StartLtEnd(t *testing.T) {
	r, err := NewNumRange(1, 2)
	assert.Nil(t, err)
	assert.Equal(t, 1, r.Start())
	assert.Equal(t, 2, r.End())
}

func TestNewNumRange_StartEqEnd(t *testing.T) {
	r, err := NewNumRange(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, 1, r.Start())
	assert.Equal(t, 1, r.End())
}

func TestNewNumRange_Include(t *testing.T) {
	r, err := NewNumRange(1, 10)
	assert.Nil(t, err)
	assert.True(t, r.include(1))
	assert.True(t, r.include(10))
	assert.True(t, r.include(5))
	assert.False(t, r.include(0))
	assert.False(t, r.include(11))
}

func TestNewNumRange_IncludeR(t *testing.T) {
	r, err := NewNumRange(1, 10)
	assert.Nil(t, err)
	r2, _ := NewNumRange(1, 10)
	assert.True(t, r.includeRange(r2))
	r2, _ = NewNumRange(1, 1)
	assert.True(t, r.includeRange(r2))
	r2, _ = NewNumRange(10, 10)
	assert.True(t, r.includeRange(r2))
	r2, _ = NewNumRange(2, 9)
	assert.True(t, r.includeRange(r2))

	r2, _ = NewNumRange(0, 10)
	assert.False(t, r.includeRange(r2))
	r2, _ = NewNumRange(1, 11)
	assert.False(t, r.includeRange(r2))
	r2, _ = NewNumRange(0, 0)
	assert.False(t, r.includeRange(r2))
	r2, _ = NewNumRange(11, 11)
	assert.False(t, r.includeRange(r2))
}

func TestNumRange_String(t *testing.T) {
	r := NumRange{
		start: 1,
		end:   10,
	}
	assert.Equal(t, "[1, 10]", r.String())
}
