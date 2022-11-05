package galois

import (
	"math"
	"reflect"
)

// IntLike is a type constraint supporting any type that reduces to an integer-like type.
type IntLike interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func getMaxTypeValue[T IntLike]() uint64 {
	var n T
	switch reflect.TypeOf(n).Kind() {
	case reflect.Int:
		return math.MaxInt
	case reflect.Int8:
		return math.MaxInt8
	case reflect.Int16:
		return math.MaxInt16
	case reflect.Int32:
		return math.MaxInt32
	case reflect.Int64:
		return math.MaxInt64

	case reflect.Uint:
		return math.MaxUint
	case reflect.Uint8:
		return math.MaxUint8
	case reflect.Uint16:
		return math.MaxUint16
	case reflect.Uint32:
		return math.MaxUint32
	case reflect.Uint64:
		return math.MaxUint64
	}
	return 0
}
