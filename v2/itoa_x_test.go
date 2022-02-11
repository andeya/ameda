package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatUintByDict(t *testing.T) {
	dict := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := uint64(0); i < 100; i++ {
		numStr := FormatUintByDict(dict, i)
		t.Logf("i=%d, s=%s", i, numStr)
		i2, err := ParseUintByDict[uint64](dict, numStr)
		assert.NoError(t, err)
		assert.Equal(t, i, i2)
	}
}
