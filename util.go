package domainerr

import (
	"fmt"
	"reflect"
)

func CheckArgument(expected bool, failureFmt string, args ...any) error {
	if expected {
		return nil
	}
	msgFmt := fmt.Sprintf("illegal argument: %s", failureFmt)
	if len(args) == 0 {
		return fmt.Errorf(msgFmt)
	}
	return fmt.Errorf(msgFmt, args...)
}

func FloatToInt(f float64) int {
	return int(f)
}

func IsNil(a any) bool {
	if a == nil {
		return true
	}

	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func NotNil(a any) bool {
	return !IsNil(a)
}
