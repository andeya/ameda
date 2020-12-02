package ameda

import (
	"reflect"
	"runtime"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestCheckGoVersion(t *testing.T) {
	defer func() { errValueUsable = nil }()
	var goVer string
	goVer, errValueUsable = checkGoVersion(runtime.Version())
	assert.NoError(t, errValueUsable)
	t.Logf("raw=%s, pure=%s", runtime.Version(), goVer)

	goVer, errValueUsable = checkGoVersion("go1.15")
	assert.NoError(t, errValueUsable)
	assert.Equal(t, "1.15", goVer)
	t.Logf("raw=%s, pure=%s", "go1.15", goVer)

	goVer, errValueUsable = checkGoVersion("go1.15rc1")
	assert.Equal(t, "1.15", goVer)
	assert.NoError(t, errValueUsable)
	t.Logf("raw=%s, pure=%s", "go1.15rc1", goVer)

	goVer, errValueUsable = checkGoVersion("go2.15rc1")
	assert.Equal(t, "2.15", goVer)
	assert.EqualError(t, errValueUsable, "required 1.9≤go<2.0, but current version is go2.15")
	t.Logf("raw=%s, pure=%s", "go2.15rc1", goVer)
}

func TestRuntimeTypeID(t *testing.T) {
	type (
		GoTime = time.Time
		Time   time.Time
		I2     interface {
			String() string
		}
		I1 interface {
			UnixNano() int64
			I2
		}
	)
	t0 := new(time.Time)
	t1 := ValueOf(t0).RuntimeTypeID()
	t2 := ValueOf(new(GoTime)).RuntimeTypeID()
	t3 := ValueOf(new(Time)).RuntimeTypeID()
	t.Log(t1, t2, t3)
	e0 := time.Time{}
	e1 := ValueOf(e0).RuntimeTypeID()
	e2 := ValueOf(GoTime{}).RuntimeTypeID()
	e3 := ValueOf(Time{}).RuntimeTypeID()
	i := ValueOf(I2(I1(&GoTime{}))).RuntimeTypeID()
	if t1 != t2 || t1 != e1 || t1 != e2 || t1 != i || t3 != e3 {
		t.FailNow()
	}
	t.Log(e1, e2, e3, i, RuntimeTypeID(reflect.TypeOf(t0)), ValueOf(t0.String).RuntimeTypeID())
}

func TestRuntimeTypeIDOf(t *testing.T) {
	type T1 struct {
		_ int
	}
	tid := RuntimeTypeIDOf(new(T1))
	t.Log(tid)
	assert.Equal(t, RuntimeTypeID(reflect.TypeOf(new(T1))), tid)
	tid2 := RuntimeTypeIDOf(T1{})
	assert.Equal(t, tid, tid2)
}

func TestKind(t *testing.T) {
	type X struct {
		A int16
		B string
	}
	var x X
	if ValueOf(&x).Kind() != reflect.Ptr {
		t.FailNow()
	}

	if ValueOf(&x).UnderlyingElem().Kind() != reflect.Struct {
		t.FailNow()
	}

	if ValueOf(x).Kind() != reflect.Struct {
		t.FailNow()
	}
	if ValueOf(x).UnderlyingElem().Kind() != reflect.Struct {
		t.FailNow()
	}

	f := func() {}
	if ValueOf(f).Kind() != reflect.Func {
		t.FailNow()
	}

	if ValueOf(t.Name).Kind() != reflect.Func {
		t.FailNow()
	}
}

func TestPointer(t *testing.T) {
	type X struct {
		A int16
		B string
	}
	x := X{A: 12345, B: "test"}
	if ValueOf(&x).Pointer() != reflect.ValueOf(&x).Pointer() {
		t.FailNow()
	}
	elemPtr := ValueOf(x).Pointer()
	a := *(*int16)(unsafe.Pointer(elemPtr))
	if a != x.A {
		t.FailNow()
	}
	b := *(*string)(unsafe.Pointer(elemPtr + unsafe.Offsetof(x.B)))
	if b != x.B {
		t.FailNow()
	}

	s := []string{""}
	if ValueOf(s).Pointer() != reflect.ValueOf(s).Pointer() {
		t.FailNow()
	}

	f := func() bool { return true }
	prt := ValueOf(f).Pointer()
	f = *(*func() bool)(unsafe.Pointer(&prt))
	if !f() {
		t.FailNow()
	}
	t.Log(ValueOf(f).FuncForPC().Name())
	prt = ValueOf(t.Name).Pointer()
	tName := *(*func() string)(unsafe.Pointer(&prt))
	if tName() != "TestPointer" {
		t.FailNow()
	}
	t.Log(ValueOf(t.Name).FuncForPC().Name())
	t.Log(ValueOf(s).FuncForPC() == nil)

}

func TestElem(t *testing.T) {
	type I interface{}
	var i I
	u := ValueFrom(reflect.ValueOf(i))
	u.Elem()

	type X struct {
		A int16
		B string
	}
	x := &X{A: 12345, B: "test"}
	xx := &x
	var elemPtr uintptr
	for i, v := range []interface{}{&xx, xx, x, *x} {
		if i == 0 {
			elemPtr = ValueOf(v).UnderlyingElem().Pointer()
		} else {
			elemPtr = ValueOf(v).Elem().Pointer()
		}
		a := *(*int16)(unsafe.Pointer(elemPtr))
		if a != x.A {
			t.FailNow()
		}
		b := *(*string)(unsafe.Pointer(elemPtr + unsafe.Offsetof(x.B)))
		if b != x.B {
			t.FailNow()
		}
	}

	var y *X
	u = ValueOf(&y)
	if u.IsNil() {
		t.FailNow()
	}
	u = u.UnderlyingElem()
	if u.Kind() != reflect.Struct {
		t.FailNow()
	}
	if !u.IsNil() {
		t.FailNow()
	}
}

func TestEmptyStruct(t *testing.T) {
	type P1 struct {
		A *int
	}
	u := ValueOf(P1{})
	if u.Pointer() != 0 {
		t.FailNow()
	}
	if !u.IsNil() {
		t.FailNow()
	}

	type P2 struct {
		A *int
		B *int
	}
	u = ValueOf(P2{})
	if u.Pointer() == 0 {
		t.FailNow()
	}
	if u.IsNil() {
		t.FailNow()
	}
}

func TestValueFrom(t *testing.T) {
	type X struct {
		A int16
		B string
	}
	x := &X{A: 12345, B: "test"}
	v := reflect.ValueOf(&x)
	u := ValueFrom2(&v).Elem()
	v = v.Elem()
	if u.RuntimeTypeID() != RuntimeTypeID(v.Type()) {
		t.FailNow()
	}
	elemPtr := u.Pointer()
	a := *(*int16)(unsafe.Pointer(elemPtr))
	if a != x.A {
		t.FailNow()
	}
	b := *(*string)(unsafe.Pointer(elemPtr + unsafe.Offsetof(x.B)))
	if b != x.B {
		t.FailNow()
	}
	if u.Pointer() != reflect.ValueOf(x).Pointer() {
		t.FailNow()
	}
}

func Benchmark_ameda(b *testing.B) {
	type T struct {
		a int
	}
	t := new(T)
	b.ReportAllocs()
	b.ResetTimer()
	u := ValueOf(t).Elem()
	for i := 0; i < b.N; i++ {
		_ = u.RuntimeTypeID()
	}
}

func Benchmark_reflect(b *testing.B) {
	type T struct {
		a int
	}
	t := new(T)
	b.ReportAllocs()
	b.ResetTimer()
	u := reflect.TypeOf(t).Elem()
	for i := 0; i < b.N; i++ {
		_ = u.String()
	}
}
