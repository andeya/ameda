package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitToDigitPtr(t *testing.T) {
	type I int64
	r, err := ToPtrWithErr(DigitToDigit[I, int16](8888))
	assert.NoError(t, err)
	assert.Equal(t, int16(8888), *r)
}

func TestToStringPtr(t *testing.T) {
	s := ToString[string](true)
	t.Logf("%T, %v", s, s)
}

func TestBoolToDigitPtr(t *testing.T) {
	s := ToPtr(BoolToDigit[bool, int16](true))
	t.Logf("%d, %v", s, *s)
}

func TestToBool(t *testing.T) {
	assert.True(t, ToBool("a"))
	assert.False(t, ToBool(""))
	assert.True(t, ToBool(1.1))
	assert.False(t, ToBool(0.0))
}
