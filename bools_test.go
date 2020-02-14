package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolsDistinct(t *testing.T) {
	a := []bool{true, true, false, true, false}
	b := BoolsDistinct(&a)
	assert.Equal(t, []bool{true, false}, a)
	assert.Equal(t, len(a), b)
}
