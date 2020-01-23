package ameda

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitPointer(t *testing.T) {
	type T struct {
		A string
		B int
		*T
	}
	var i ****T
	var i2 interface{} = &i
	var i3 = &i2
	var i4 = &i3
	var i5 interface{} = &i4
	var i6 = &i5
	var i7 interface{} = &i6
	v := reflect.ValueOf(&i7)
	done := InitPointer(v)
	assert.True(t, done)
	assert.Equal(t, T{}, ****i)
}
