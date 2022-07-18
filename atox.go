package atox

import (
	"fmt"
	"reflect"
	"strconv"

	"golang.org/x/exp/constraints"
)

func Must[T constraints.Float | constraints.Integer](s string) T {
	t, err := N[T](s)
	if err != nil {
		panic(err)
	}
	return t
}

func N[T constraints.Float | constraints.Integer](s string) (T, error) {
	var z T
	rt := reflect.TypeOf(z)
	switch rt.Kind() {
	case reflect.Float32, reflect.Float64:
		t, err := strconv.ParseFloat(s, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		t, err := strconv.ParseInt(s, 10, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		t, err := strconv.ParseUint(s, 10, rt.Bits())
		if err != nil {
			return z, err
		}
		return T(t), nil
	default:
		return z, fmt.Errorf("impossible?")
	}
}
