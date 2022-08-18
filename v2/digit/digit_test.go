package digit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigit(t *testing.T) {
	assert.Equal(t, 1, Abs(-1))
}
