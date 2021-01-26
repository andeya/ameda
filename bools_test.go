package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolsDistinct(t *testing.T) {
	a := []bool{true, true, false, true, false}
	b := BoolsDistinct(&a, true)
	assert.Equal(t, []bool{true, false}, a)
	assert.Equal(t, len(a), len(b))
	assert.Equal(t, map[bool]int{false: 2, true: 3}, b)
}

func TestBoolsRemoveFirst(t *testing.T) {
	var a = []bool{true, true, false, false}
	assert.Equal(t, 3, BoolsRemoveFirst(&a, false))
	assert.Equal(t, []bool{true, true, false}, a)
}

func TestBoolsRemoveEvery(t *testing.T) {
	var a = []bool{true, true, false, false, true}
	assert.Equal(t, 3, BoolsRemoveEvery(&a, false))
	assert.Equal(t, []bool{true, true, true}, a)
}
