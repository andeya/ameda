package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBool(t *testing.T) {
	cases := []struct {
		actualParam  interface{}
		emptyAsFalse bool
		expected     bool
	}{
		{"", false, false},
		{"", true, false},
		{0, true, false},
		{nil, true, false},
		{(*testing.T)(nil), true, false},
		{(*testing.T)(nil), false, false},
		{t, false, false},
		{t, true, true},
		{struct{ int }{}, false, false},
		{struct{ int }{}, true, false},
	}
	for _, c := range cases {
		actual, err := InterfaceToBool(c.actualParam, c.emptyAsFalse)
		if c.emptyAsFalse {
			assert.NoError(t, err, c)
		} else {
			assert.NotNil(t, err, c)
		}
		assert.Equal(t, c.expected, actual, c)
	}
}
