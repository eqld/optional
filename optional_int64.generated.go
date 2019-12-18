package optional

// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT

import (
	"encoding/json"
	"errors"
)

// Int64 is a superset of type int64 that has explicit "empty" value.
type Int64 struct {
	Value   int64
	Present bool
}

// MakeInt64 converts int64 to Int64.
func MakeInt64(value int64) Int64 {
	return Int64{Value: value, Present: true}
}

// MakeInt64FromPtr converts a pointer to int64 to Int64.
func MakeInt64FromPtr(ptr *int64) Int64 {
	if ptr == nil {
		return Int64{}
	}

	return MakeInt64(*ptr)
}

// SafeValue safely converts Int64 to int64 returning its default value if the value of Int64 is not present.
func (o Int64) SafeValue() (value int64) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts Int64 to a pointer to int64 returning nil pointer if the value of Int64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Int64.
func (o Int64) SafePtr() (ptr *int64) {
	if o.Present {
		ptr = new(int64)
		*ptr = o.Value
	}

	return
}

// SafeValueWithFlag safely converts Int64 to int64 returning it with a flag set to true
// or its default value and the flag set to false if the value of Int64 is not present.
func (o Int64) SafeValueWithFlag() (value int64, ok bool) {
	if !o.Present {
		return
	}

	return o.Value, o.Present
}

// SafeValueWithErr safely converts Int64 to int64 returning it with nil error
// or its default value and non-nil error if the value of Int64 is not present.
func (o Int64) SafeValueWithErr() (value int64, err error) {
	if !o.Present {
		err = errors.New("value of optional.Int64 is not present")
		return
	}

	return o.Value, nil
}

// SafePtrWithFlag safely converts Int64 to a pointer to int64 returning a pointer to a value with a flag set to true
// or nil pointer and the flag set to false if the value of Int64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Int64.
func (o Int64) SafePtrWithFlag() (ptr *int64, ok bool) {
	if !o.Present {
		return
	}

	ptr = new(int64)
	*ptr, ok = o.Value, o.Present

	return
}

// SafePtrWithErr safely converts Int64 to a pointer to int64 returning a pointer to a value with nil error
// or nil pointer and non-nil error if the value of Int64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Int64.
func (o Int64) SafePtrWithErr() (ptr *int64, err error) {
	if !o.Present {
		err = errors.New("value of optional.Int64 is not present")
		return
	}

	ptr = new(int64)
	*ptr = o.Value

	return
}

// Equals returns true if both values are not present or both values are present and are equal according to a provided determinant.
func (o Int64) Equals(other Int64, determinant func(this, other int64) bool) bool {
	return (!o.Present && !other.Present) || (o.Present && other.Present && determinant(o.Value, other.Value))
}

// Compare returns a result of provided comparator and true if both values are present, otherwise it returns 0 and false.
func (o Int64) Compare(other Int64, comparator func(this, other int64) int) (int, bool) {
	if !o.Present || !other.Present {
		return 0, false
	}

	return comparator(o.Value, other.Value), true
}

// Filter returns the Int64 if its value is present and it matches the given predicate, otherwise it returns an empty Int64.
func (o Int64) Filter(test func(value int64) bool) (result Int64) {
	if !o.Present || !test(o.Value) {
		return
	}

	return o
}

// Map applies the provided mapping function to a value and returns its result as Int64 if the value is present,
// otherwise is returns an empty Int64.
func (o Int64) Map(mapper func(value int64) (result int64, present bool)) (result Int64) {
	if !o.Present {
		return
	}

	result.Value, result.Present = mapper(o.Value)

	return
}

// IfPresent invokes the specified action with the value if it is present.
func (o Int64) IfPresent(action func(value int64)) {
	if o.Present {
		action(o.Value)
	}
}

// OrElse returns the value if it is present, otherwise it returns given other value.
func (o Int64) OrElse(other int64) int64 {
	if !o.Present {
		return other
	}

	return o.Value
}

// OrElseWithFlag returns the value if it is present with a flag set to true, otherwise it returns given other value
// and the flag set to false.
func (o Int64) OrElseWithFlag(other int64) (int64, bool) {
	if !o.Present {
		return other, false
	}

	return o.Value, o.Present
}

// OrElseWithErr returns the value if it is present with nil error, otherwise it returns given other value
// and non-nil error.
func (o Int64) OrElseWithErr(other int64) (int64, error) {
	if !o.Present {
		return other, errors.New("value of optional.Int64 is not present")
	}

	return o.Value, nil
}

// OrElseGet returns the value if it is present, otherwise it invokes a supplier and returns a result of that invocation.
func (o Int64) OrElseGet(supplier func() int64) int64 {
	if !o.Present {
		return supplier()
	}

	return o.Value
}

// OrElseGetWithFlag returns the value if it is present with a flag set to true, otherwise it invokes a supplier and returns
// a result of that invocation with a flag set to false.
func (o Int64) OrElseGetWithFlag(supplier func() int64) (result int64, ok bool) {
	if !o.Present {
		return supplier(), false
	}

	return o.Value, o.Present
}

// OrElseGetWithErr returns the value if it is present with nil error, otherwise it invokes a supplier and returns
// a result of that invocation with non-nil error.
func (o Int64) OrElseGetWithErr(supplier func() int64) (result int64, err error) {
	if !o.Present {
		return supplier(), errors.New("value of optional.Int64 is not present")
	}

	return o.Value, nil
}

// OrElseErr returns the value if it is present with nil error, otherwise it invokes an error supplier and returns
// default value and an error returned by the error supplier.
func (o Int64) OrElseErr(errSupplier func() error) (result int64, err error) {
	if !o.Present {
		err = errSupplier()
		return
	}

	return o.Value, nil
}

// MarshalJSON marshals Int64 to json.
func (o Int64) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Int64 from json.
func (o *Int64) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
