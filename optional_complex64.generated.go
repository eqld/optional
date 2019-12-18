package optional

// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT

import (
	"encoding/json"
	"errors"
)

// Complex64 is a superset of type complex64 that has explicit "empty" value.
type Complex64 struct {
	Value   complex64
	Present bool
}

// MakeComplex64 converts complex64 to Complex64.
func MakeComplex64(value complex64) Complex64 {
	return Complex64{Value: value, Present: true}
}

// MakeComplex64FromPtr converts a pointer to complex64 to Complex64.
func MakeComplex64FromPtr(ptr *complex64) Complex64 {
	if ptr == nil {
		return Complex64{}
	}

	return MakeComplex64(*ptr)
}

// SafeValue safely converts Complex64 to complex64 returning its default value if the value of Complex64 is not present.
func (o Complex64) SafeValue() (value complex64) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts Complex64 to a pointer to complex64 returning nil pointer if the value of Complex64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Complex64.
func (o Complex64) SafePtr() (ptr *complex64) {
	if o.Present {
		ptr = new(complex64)
		*ptr = o.Value
	}

	return
}

// SafeValueWithFlag safely converts Complex64 to complex64 returning it with a flag set to true
// or its default value and the flag set to false if the value of Complex64 is not present.
func (o Complex64) SafeValueWithFlag() (value complex64, ok bool) {
	if !o.Present {
		return
	}

	return o.Value, o.Present
}

// SafeValueWithErr safely converts Complex64 to complex64 returning it with nil error
// or its default value and non-nil error if the value of Complex64 is not present.
func (o Complex64) SafeValueWithErr() (value complex64, err error) {
	if !o.Present {
		err = errors.New("value of optional.Complex64 is not present")
		return
	}

	return o.Value, nil
}

// SafePtrWithFlag safely converts Complex64 to a pointer to complex64 returning a pointer to a value with a flag set to true
// or nil pointer and the flag set to false if the value of Complex64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Complex64.
func (o Complex64) SafePtrWithFlag() (ptr *complex64, ok bool) {
	if !o.Present {
		return
	}

	ptr = new(complex64)
	*ptr, ok = o.Value, o.Present

	return
}

// SafePtrWithErr safely converts Complex64 to a pointer to complex64 returning a pointer to a value with nil error
// or nil pointer and non-nil error if the value of Complex64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Complex64.
func (o Complex64) SafePtrWithErr() (ptr *complex64, err error) {
	if !o.Present {
		err = errors.New("value of optional.Complex64 is not present")
		return
	}

	ptr = new(complex64)
	*ptr = o.Value

	return
}

// Equals returns true if both values are not present or both values are present and are equal according to a provided determinant.
func (o Complex64) Equals(other Complex64, determinant func(this, other complex64) bool) bool {
	return (!o.Present && !other.Present) || (o.Present && other.Present && determinant(o.Value, other.Value))
}

// Compare returns a result of provided comparator and true if both values are present, otherwise it returns 0 and false.
func (o Complex64) Compare(other Complex64, comparator func(this, other complex64) int) (int, bool) {
	if !o.Present || !other.Present {
		return 0, false
	}

	return comparator(o.Value, other.Value), true
}

// Filter returns the Complex64 if its value is present and it matches the given predicate, otherwise it returns an empty Complex64.
func (o Complex64) Filter(test func(value complex64) bool) (result Complex64) {
	if !o.Present || !test(o.Value) {
		return
	}

	return o
}

// Map applies the provided mapping function to a value and returns its result as Complex64 if the value is present,
// otherwise is returns an empty Complex64.
func (o Complex64) Map(mapper func(value complex64) (result complex64, present bool)) (result Complex64) {
	if !o.Present {
		return
	}

	result.Value, result.Present = mapper(o.Value)

	return
}

// IfPresent invokes the specified action with the value if it is present.
func (o Complex64) IfPresent(action func(value complex64)) {
	if o.Present {
		action(o.Value)
	}
}

// OrElse returns the value if it is present, otherwise it returns given other value.
func (o Complex64) OrElse(other complex64) complex64 {
	if !o.Present {
		return other
	}

	return o.Value
}

// OrElseWithFlag returns the value if it is present with a flag set to true, otherwise it returns given other value
// and the flag set to false.
func (o Complex64) OrElseWithFlag(other complex64) (complex64, bool) {
	if !o.Present {
		return other, false
	}

	return o.Value, o.Present
}

// OrElseWithErr returns the value if it is present with nil error, otherwise it returns given other value
// and non-nil error.
func (o Complex64) OrElseWithErr(other complex64) (complex64, error) {
	if !o.Present {
		return other, errors.New("value of optional.Complex64 is not present")
	}

	return o.Value, nil
}

// OrElseGet returns the value if it is present, otherwise it invokes a supplier and returns a result of that invocation.
func (o Complex64) OrElseGet(supplier func() complex64) complex64 {
	if !o.Present {
		return supplier()
	}

	return o.Value
}

// OrElseGetWithFlag returns the value if it is present with a flag set to true, otherwise it invokes a supplier and returns
// a result of that invocation with a flag set to false.
func (o Complex64) OrElseGetWithFlag(supplier func() complex64) (result complex64, ok bool) {
	if !o.Present {
		return supplier(), false
	}

	return o.Value, o.Present
}

// OrElseGetWithErr returns the value if it is present with nil error, otherwise it invokes a supplier and returns
// a result of that invocation with non-nil error.
func (o Complex64) OrElseGetWithErr(supplier func() complex64) (result complex64, err error) {
	if !o.Present {
		return supplier(), errors.New("value of optional.Complex64 is not present")
	}

	return o.Value, nil
}

// OrElseErr returns the value if it is present with nil error, otherwise it invokes an error supplier and returns
// default value and an error returned by the error supplier.
func (o Complex64) OrElseErr(errSupplier func() error) (result complex64, err error) {
	if !o.Present {
		err = errSupplier()
		return
	}

	return o.Value, nil
}

// MarshalJSON marshals Complex64 to json.
func (o Complex64) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Complex64 from json.
func (o *Complex64) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
