package ameda

import (
	"encoding/json"
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

func TestInitSampleValue(t *testing.T) {
	type P3 struct {
		M string
	}
	type P2 struct {
		X float32
		Y uint8
		Z **bool
	}
	type P struct {
		A  string
		B  int
		Po P2
		Ps []P2
		P2
		P3
		Pm map[string]P
	}
	v := InitSampleValue(reflect.TypeOf(map[string]P{}), 5)
	b, err := json.MarshalIndent(v.Interface(), "", "  ")
	assert.NoError(t, err)
	t.Logf("%s", b)
}
