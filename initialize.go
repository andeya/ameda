package ameda

import "reflect"

// InitAndGetString if strPtr is empty string, initialize it with def,
// and return the final value.
func InitAndGetString(strPtr *string, def string) string {
	if strPtr == nil {
		return def
	}
	if *strPtr == "" {
		*strPtr = def
	}
	return *strPtr
}

// InitPointer initializes null pointer.
func InitPointer(v reflect.Value) bool {
	for {
		kind := v.Kind()
		if kind == reflect.Interface {
			v = v.Elem()
			continue
		}
		if kind != reflect.Ptr {
			return true
		}
		u := v.Elem()
		if u.IsValid() {
			v = u
			continue
		}
		if !v.CanSet() {
			return false
		}
		v2 := reflect.New(v.Type().Elem())
		v.Set(v2)
		v = v.Elem()
	}
}
