package ameda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUintByDict(t *testing.T) {
	dict := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numStr := "DDEZQ"
	num, err := ParseUintByDict[uint64](dict, numStr)
	assert.NoError(t, err)
	t.Logf("DDEZQ=%d", num) // DDEZQ=1427026
	numStr2 := FormatUintByDict(dict, num)
	assert.Equal(t, numStr2, numStr)
}
