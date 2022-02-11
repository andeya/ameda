// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ameda

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	BitSizeError = bitSizeError
	BaseError    = baseError
)

func TestStrconv62(t *testing.T) {
	a := FormatInt(math.MaxInt64, 62)
	// FormatInt: base10=9223372036854775807 -> base62=aZl8N0y58M7
	t.Logf("FormatInt: base10=%d -> base62=%s", math.MaxInt64, a)
	i, err := ParseInt(a, 62, 64)
	assert.NoError(t, err)
	assert.Equal(t, int64(math.MaxInt64), i)
	t.Logf("ParseInt base62: bitSize=64, num=%d", i)
	i, err = ParseInt(a, 62, 32)
	assert.EqualError(t, err.(*strconv.NumError).Err, strconv.ErrRange.Error())
	assert.Equal(t, int64(math.MaxInt32), i)

	a = FormatInt(math.MinInt64, 62)
	// FormatInt: base10=-9223372036854775808 -> base62=-aZl8N0y58M8
	t.Logf("FormatInt: base10=%d -> base62=%s", math.MinInt64, a)
	i, err = ParseInt(a, 62, 64)
	assert.NoError(t, err)
	assert.Equal(t, int64(math.MinInt64), i)
	t.Logf("ParseInt base62: bitSize=64, num=%d", i)
	i, err = ParseInt(a, 62, 32)
	assert.EqualError(t, err.(*strconv.NumError).Err, strconv.ErrRange.Error())
	assert.Equal(t, int64(math.MinInt32), i)
}
