package assert

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	compareTypeEqual   uint = 1
	compareTypeGreater      = 1 << 1
	compareTypeLess         = 1 << 2
)

const (
	compareValueTypeInt int = iota
	compareValueTypeUint
	compareValueTypeFloat
	compareValueTypeString
)

// Gt compares the values and sets the result to false if the first value is not greater than to
// the second value.
//
//	a.Gt(2, 1) // success
//	a.Gt(3.14, 1.68) // success
//	a.Gt("BCD", "ABC") // success
//	a.Gt(2, 2) // fail
//	a.Gt(1, 2) // fail
func (a *Assertion) Gt(v1, v2 any, message ...string) error {
	a.T.Helper()

	return tryCompareOrderableValues(
		a.T,
		false,
		compareTypeGreater,
		v1, v2,
		fmt.Sprintf(defaultErrMessageGt, v1, v2),
		message...,
	)
}

// GtNow compares the values and sets the result to false if the first value is not greater than to
// the second value. It will panic if they do not match the expected result.
//
//	a.GtNow(2, 1) // success
//	a.GtNow(3.14, 1.68) // success
//	a.GtNow("BCD", "ABC") // success
//	a.GtNow(1, 2) // fail and terminate
//	// never runs
func (a *Assertion) GtNow(v1, v2 any, message ...string) error {
	a.T.Helper()

	return tryCompareOrderableValues(
		a.T,
		true,
		compareTypeGreater,
		v1, v2,
		fmt.Sprintf(defaultErrMessageGt, v1, v2),
		message...,
	)
}

// Gte compares the values and sets the result to false if the first value is not greater than or
// equal to the second value.
//
//	a.Gte(2, 1) // success
//	a.Gte(3.14, 1.68) // success
//	a.Gte("BCD", "ABC") // success
//	a.Gte(2, 2) // success
//	a.Gte(1, 2) // fail
func (a *Assertion) Gte(v1, v2 any, message ...string) error {
	a.T.Helper()

	return tryCompareOrderableValues(
		a.T,
		false,
		compareTypeEqual|compareTypeGreater,
		v1, v2,
		fmt.Sprintf(defaultErrMessageGte, v1, v2),
		message...,
	)
}

// GteNow compares the values and sets the result to false if the first value is not greater than
// or equal to the second value. It will panic if they do not match the expected result.
//
//	a.GteNow(2, 1) // success
//	a.GteNow(3.14, 1.68) // success
//	a.GteNow("BCD", "ABC") // success
//	a.GteNow(2, 2) // success
//	a.GteNow(1, 2) // fail and terminate
//	// never runs
func (a *Assertion) GteNow(v1, v2 any, message ...string) error {
	a.T.Helper()

	return tryCompareOrderableValues(
		a.T,
		true,
		compareTypeEqual|compareTypeGreater,
		v1, v2,
		fmt.Sprintf(defaultErrMessageGte, v1, v2),
		message...,
	)
}

// Lt compares the values and sets the result to false if the first value is not less than the
// second value.
//
//	a.Lt(1, 2) // success
//	a.Lt(1.68, 3.14) // success
//	a.Lt("ABC", "BCD") // success
//	a.Lt(2, 2) // fail
//	a.Lt(2, 1) // fail
func (a *Assertion) Lt(v1, v2 any, message ...string) error {
	a.T.Helper()

	return tryCompareOrderableValues(
		a.T,
		false,
		compareTypeLess,
		v1, v2,
		fmt.Sprintf(defaultErrMessageLt, v1, v2),
		message...,
	)
}

// LtNow compares the values and sets the result to false if the first value is not less than the
// second value. It will panic if they do not match the expected result.
//
//	a.LtNow(1, 2) // success
//	a.LtNow(1.68, 3.14) // success
//	a.LtNow("ABC", "BCD") // success
//	a.LtNow(2, 1) // fail and terminate
//	// never runs
func (a *Assertion) LtNow(v1, v2 any, message ...string) error {
	a.T.Helper()

	return tryCompareOrderableValues(
		a.T,
		true,
		compareTypeLess,
		v1, v2,
		fmt.Sprintf(defaultErrMessageLt, v1, v2),
		message...,
	)
}

// Lte compares the values and sets the result to false if the first value is not less than or
// equal to the second value.
//
//	a.Lte(1, 2) // success
//	a.Lte(1.68, 3.14) // success
//	a.Lte("ABC", "BCD") // success
//	a.Lte(2, 2) // success
//	a.Lte(2, 1) // fail
func (a *Assertion) Lte(v1, v2 any, message ...string) error {
	a.T.Helper()

	return tryCompareOrderableValues(
		a.T,
		false,
		compareTypeEqual|compareTypeLess,
		v1, v2,
		fmt.Sprintf(defaultErrMessageLte, v1, v2),
		message...,
	)
}

// LteNow compares the values and sets the result to false if the first value is not less than or
// equal to the second value. It will panic if they do not match the expected result.
//
//	a.LteNow(1, 2) // success
//	a.LteNow(1.68, 3.14) // success
//	a.LteNow("ABC", "BCD") // success
//	a.LteNow(2, 2) // success
//	a.LteNow(2, 1) // fail and terminate
//	// never runs
func (a *Assertion) LteNow(v1, v2 any, message ...string) error {
	a.T.Helper()

	return tryCompareOrderableValues(
		a.T,
		true,
		compareTypeEqual|compareTypeLess,
		v1, v2,
		fmt.Sprintf(defaultErrMessageLte, v1, v2),
		message...,
	)
}

// tryCompareOrderableValues tries to compare the values by the comparison type, and returns an
// error if the result does not match.
func tryCompareOrderableValues(
	t *testing.T,
	failedNow bool,
	compareType uint,
	v1, v2 any,
	defaultMessage string,
	message ...string,
) error {
	t.Helper()

	vv1 := reflect.ValueOf(v1)
	vv2 := reflect.ValueOf(v2)

	if !isSameType(vv1.Type(), vv2.Type()) {
		panic(ErrNotSameType)
	} else if !isOrderable(v1) {
		panic(ErrNotOrderable)
	}

	return test(
		t,
		func() bool { return compareValues(vv1, vv2, compareType) },
		failedNow,
		defaultMessage,
		message,
	)
}

// compareValues tries to compare the values by the comparison type.
func compareValues(v1, v2 reflect.Value, compareType uint) bool {
	k := v1.Kind()
	t := 0 // 1: int, 2: uint, 3: float, 4: string
	var i1, i2 int64
	var u1, u2 uint64
	var f1, f2 float64
	var s1, s2 string

	switch {
	case k >= reflect.Int && k <= reflect.Int64:
		t = compareValueTypeInt
		i1 = v1.Int()
		i2 = v2.Int()
	case k >= reflect.Uint && k <= reflect.Uintptr:
		t = compareValueTypeUint
		u1 = v1.Uint()
		u2 = v2.Uint()
	case k == reflect.Float32 || k == reflect.Float64:
		t = compareValueTypeFloat
		f1 = v1.Float()
		f2 = v2.Float()
	case k == reflect.String:
		t = compareValueTypeString
		s1 = v1.String()
		s2 = v2.String()
	default:
		return false
	}

	if (compareType & compareTypeEqual) > 0 {
		if (t == compareValueTypeInt && i1 == i2) ||
			(t == compareValueTypeUint && u1 == u2) ||
			(t == compareValueTypeFloat && f1 == f2) ||
			(t == compareValueTypeString && s1 == s2) {
			return true
		}
	}
	if (compareType & compareTypeGreater) > 0 {
		if (t == compareValueTypeInt && i1 > i2) ||
			(t == compareValueTypeUint && u1 > u2) ||
			(t == compareValueTypeFloat && f1 > f2) ||
			(t == compareValueTypeString && s1 > s2) {
			return true
		}
	}
	if (compareType & compareTypeLess) > 0 {
		if (t == compareValueTypeInt && i1 < i2) ||
			(t == compareValueTypeUint && u1 < u2) ||
			(t == compareValueTypeFloat && f1 < f2) ||
			(t == compareValueTypeString && s1 < s2) {
			return true
		}
	}

	return false
}

// isOrderable gets the type of the value, and checks whether the type is comparable or not.
func isOrderable(v any) bool {
	switch v.(type) {
	case
		int, int8, int16, int32, int64, // Signed integer
		uint, uint8, uint16, uint32, uint64, uintptr, // Unsigned integer
		float32, float64, // Floating-point number
		string: // string
		return true
	default:
		kind := reflect.TypeOf(v).Kind()
		return (kind >= reflect.Int && kind <= reflect.Int64) ||
			(kind >= reflect.Uint && kind <= reflect.Uintptr) ||
			(kind >= reflect.Float32 && kind <= reflect.Float64) ||
			kind == reflect.String
	}
}
