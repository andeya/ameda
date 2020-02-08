package ameda

import (
	"fmt"
	"reflect"
)

// InterfaceToInterfacePtr converts interface to *interface.
func InterfaceToInterfacePtr(i interface{}) *interface{} {
	return &i
}

// InterfaceToString converts interface to string.
func InterfaceToString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

// InterfaceToStringPtr converts interface to *string.
func InterfaceToStringPtr(i interface{}) *string {
	v := InterfaceToString(i)
	return &v
}

// InterfaceToBool converts interface to bool.
// NOTE:
//  0 is false, other numbers are true
func InterfaceToBool(i interface{}, emptyAsFalse ...bool) (bool, error) {
	switch v := i.(type) {
	case bool:
		return v, nil
	case nil:
		return false, nil
	case float32:
		return Float32ToBool(v), nil
	case float64:
		return Float64ToBool(v), nil
	case int:
		return IntToBool(v), nil
	case int8:
		return Int8ToBool(v), nil
	case int16:
		return Int16ToBool(v), nil
	case int32:
		return Int32ToBool(v), nil
	case int64:
		return Int64ToBool(v), nil
	case uint:
		return UintToBool(v), nil
	case uint8:
		return Uint8ToBool(v), nil
	case uint16:
		return Uint16ToBool(v), nil
	case uint32:
		return Uint32ToBool(v), nil
	case uint64:
		return Uint64ToBool(v), nil
	case uintptr:
		return v != 0, nil
	case string:
		return StringToBool(v, emptyAsFalse...)
	default:
		r := IndirectValue(reflect.ValueOf(i))
		switch r.Kind() {
		case reflect.Bool:
			return r.Bool(), nil
		case reflect.Invalid:
			return false, nil
		case reflect.Float32, reflect.Float64:
			return Float64ToBool(r.Float()), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return Int64ToBool(r.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return Uint64ToBool(r.Uint()), nil
		case reflect.String:
			return StringToBool(r.String(), emptyAsFalse...)
		}
		if isEmptyAsZero(emptyAsFalse) {
			return !r.IsZero(), nil
		}
		return false, fmt.Errorf("cannot convert %#v of type %T to bool", i, i)
	}
}

// InterfaceToBoolPtr converts interface to *bool.
// NOTE:
//  0 is false, other numbers are true
func InterfaceToBoolPtr(i interface{}, emptyAsFalse ...bool) (*bool, error) {
	r, err := InterfaceToBool(i, emptyAsFalse...)
	return &r, err
}

// InterfaceToFloat32 converts interface to float32.
func InterfaceToFloat32(i interface{}, emptyStringAsZero ...bool) (float32, error) {
	switch v := i.(type) {
	case bool:
		return BoolToFloat32(v), nil
	case nil:
		return 0, nil
	case int:
		return IntToFloat32(v), nil
	case int8:
		return Int8ToFloat32(v), nil
	case int16:
		return Int16ToFloat32(v), nil
	case int32:
		return Int32ToFloat32(v), nil
	case int64:
		return Int64ToFloat32(v), nil
	case uint:
		return UintToFloat32(v), nil
	case uint8:
		return Uint8ToFloat32(v), nil
	case uint16:
		return Uint16ToFloat32(v), nil
	case uint32:
		return Uint32ToFloat32(v), nil
	case uint64:
		return Uint64ToFloat32(v), nil
	case uintptr:
		return UintToFloat32(uint(v)), nil
	case string:
		return StringToFloat32(v, emptyStringAsZero...)
	default:
		r := IndirectValue(reflect.ValueOf(i))
		switch r.Kind() {
		case reflect.Bool:
			return BoolToFloat32(r.Bool()), nil
		case reflect.Invalid:
			return 0, nil
		case reflect.Float32:
			return float32(r.Float()), nil
		case reflect.Float64:
			return Float64ToFloat32(r.Float())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return Int64ToFloat32(r.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return Uint64ToFloat32(r.Uint()), nil
		case reflect.String:
			return StringToFloat32(r.String(), emptyStringAsZero...)
		}
		if isEmptyAsZero(emptyStringAsZero) {
			return BoolToFloat32(!r.IsZero()), nil
		}
		return 0, fmt.Errorf("cannot convert %#v of type %T to float32", i, i)
	}
}

// InterfaceToFloat32Ptr converts interface to *float32.
func InterfaceToFloat32Ptr(i interface{}, emptyAsFalse ...bool) (*float32, error) {
	r, err := InterfaceToFloat32(i, emptyAsFalse...)
	return &r, err
}
