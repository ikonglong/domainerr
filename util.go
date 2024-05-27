package domainerr

import (
	"fmt"
	"reflect"
	"strings"
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

func TraceCauseOnce(err error) error {
	switch e := err.(type) {
	case interface{ Cause() error }:
		return e.Cause()
	}
	return nil
}

// ChainMsg strings each error's message in the given chain together with separator '->'.
//
// The next error in the chain is got by `interface{ Cause() error }`. If an error implements
// `interface{ Unwrap() error }`, this error wrapper and the wrapped error are together handled
// as an error on the chain.
func ChainMsg(chain error) string {
	if chain == nil {
		return ""
	}

	var err error = chain
	var sb strings.Builder
	reachEnd := false
	for {
		switch v := err.(type) {
		case nil:
			reachEnd = true
		case *Error:
			if IsNil(v) {
				reachEnd = true
				break
			}
			sb.WriteString(v.Status().Message())
		case error:
			if IsNil(v) {
				reachEnd = true
				break
			}
			sb.WriteString(v.Error())
		}

		if reachEnd {
			break
		} else {
			sb.WriteString(" -> ")
			err = TraceCauseOnce(err)
		}
	}
	s := sb.String()
	return s[:len(s)-4]
}
