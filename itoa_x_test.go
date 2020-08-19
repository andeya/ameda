package ameda

import "testing"

func TestFormatUintByDict(t *testing.T) {
	dict := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := uint64(0); i < 100; i++ {
		t.Logf("i=%d, s=%s", i, FormatUintByDict(dict, i))
	}
}
