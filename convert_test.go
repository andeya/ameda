package ameda

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsafeBytesToString(t *testing.T) {
	var b = []byte("abc")
	s := UnsafeBytesToString(b)
	assert.Equal(t, string(b), s)
}

func TestUnsafeStringToBytes(t *testing.T) {
	var s = "abc"
	b := UnsafeStringToBytes(s)
	assert.Equal(t, []byte(s), b)
}

func TestReferenceSlice(t *testing.T) {
	v := reflect.ValueOf([]int{1, 2})
	v = ReferenceSlice(v, 1)
	ret := v.Interface().([]*int)
	t.Logf("%#v", ret)

	v = reflect.ValueOf([]int{})
	v = ReferenceSlice(v, 1)
	ret = v.Interface().([]*int)
	t.Logf("%#v", ret)
}

func TestDereferenceSlice(t *testing.T) {
	one := 1
	two := 2
	v := reflect.ValueOf([]*int{&one, &two})
	v = DereferenceSlice(v)
	ret := v.Interface().([]int)
	t.Logf("%#v", ret)

	v = reflect.ValueOf([]*int{})
	v = DereferenceSlice(v)
	ret = v.Interface().([]int)
	t.Logf("%#v", ret)
}
