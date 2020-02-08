package ameda

import (
	"math"
	"strconv"
)

// UintToInterface converts int to interface.
func UintToInterface(v uint) interface{} {
	return v
}

// UintToInterfacePtr converts int to *interface.
func UintToInterfacePtr(v uint) *interface{} {
	r := UintToInterface(v)
	return &r
}

// UintToString converts int to string.
func UintToString(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

// UintToStringPtr converts int to *string.
func UintToStringPtr(v uint) *string {
	r := UintToString(v)
	return &r
}

// UintToBool converts int to bool.
func UintToBool(v uint) bool {
	return v != 0
}

// UintToBoolPtr converts int to *bool.
func UintToBoolPtr(v uint) *bool {
	r := UintToBool(v)
	return &r
}

// UintToFloat32 converts int to float32.
func UintToFloat32(v uint) float32 {
	return float32(v)
}

// UintToFloat32Ptr converts int to *float32.
func UintToFloat32Ptr(v uint) *float32 {
	r := UintToFloat32(v)
	return &r
}

// UintToFloat64 converts int to float64.
func UintToFloat64(v uint) float64 {
	return float64(v)
}

// UintToFloat64Ptr converts int to *float64.
func UintToFloat64Ptr(v uint) *float64 {
	r := UintToFloat64(v)
	return &r
}

// UintToInt converts int to int.
func UintToInt(v uint) (int, error) {
	if is64BitPlatform {
		if v > math.MaxInt64 {
			return 0, errOverflowValue
		}
	} else {
		if v > math.MaxInt32 {
			return 0, errOverflowValue
		}
	}
	return int(v), nil
}

// UintToIntPtr converts int to *int.
func UintToIntPtr(v uint) (*int, error) {
	r, err := UintToInt(v)
	return &r, err
}

// UintToInt8 converts int to int8.
func UintToInt8(v uint) (int8, error) {
	if v > math.MaxInt8 {
		return 0, errOverflowValue
	}
	return int8(v), nil
}

// UintToInt8Ptr converts int to *int8.
func UintToInt8Ptr(v uint) (*int8, error) {
	r, err := UintToInt8(v)
	return &r, err
}

// UintToInt16 converts int to int16.
func UintToInt16(v uint) (int16, error) {
	if v > math.MaxInt16 {
		return 0, errOverflowValue
	}
	return int16(v), nil
}

// UintToInt16Ptr converts int to *int16.
func UintToInt16Ptr(v uint) (*int16, error) {
	r, err := UintToInt16(v)
	return &r, err
}

// UintToInt32 converts int to int32.
func UintToInt32(v uint) (int32, error) {
	if v > math.MaxInt32 {
		return 0, errOverflowValue
	}
	return int32(v), nil
}

// UintToInt32Ptr converts int to *int32.
func UintToInt32Ptr(v uint) (*int32, error) {
	r, err := UintToInt32(v)
	return &r, err
}

// UintToInt64 converts int to int64.
func UintToInt64(v uint) (int64, error) {
	if v > math.MaxInt64 {
		return 0, errOverflowValue
	}
	return int64(v), nil
}

// UintToInt64Ptr converts int to *int64.
func UintToInt64Ptr(v uint) (*int64, error) {
	r, err := UintToInt64(v)
	return &r, err
}

// UintToUintPtr converts int to *uint.
func UintToUintPtr(v uint) *uint {
	return &v
}

// UintToUint8 converts int to uint8.
func UintToUint8(v uint) (uint8, error) {
	if v > math.MaxUint8 {
		return 0, errOverflowValue
	}
	return uint8(v), nil
}

// UintToUint8Ptr converts int to *uint8.
func UintToUint8Ptr(v uint) (*uint8, error) {
	r, err := UintToUint8(v)
	return &r, err
}

// UintToUint16 converts int to uint16.
func UintToUint16(v uint) (uint16, error) {
	if v > math.MaxUint16 {
		return 0, errOverflowValue
	}
	return uint16(v), nil
}

// UintToUint16Ptr converts int to *uint16.
func UintToUint16Ptr(v uint) (*uint16, error) {
	r, err := UintToUint16(v)
	return &r, err
}

// UintToUint32 converts int to uint32.
func UintToUint32(v uint) (uint32, error) {
	if is64BitPlatform && v > math.MaxUint32 {
		return 0, errOverflowValue
	}
	return uint32(v), nil
}

// UintToUint32Ptr converts int to *uint32.
func UintToUint32Ptr(v uint) (*uint32, error) {
	r, err := UintToUint32(v)
	return &r, err
}

// UintToUint64 converts int to uint64.
func UintToUint64(v uint) uint64 {
	return uint64(v)
}

// UintToUint64Ptr converts int to *uint64.
func UintToUint64Ptr(v uint) *uint64 {
	r := UintToUint64(v)
	return &r
}
