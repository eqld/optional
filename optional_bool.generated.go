package optional

// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT

import (
	"encoding/json"
	"errors"
)

// Bool is a superset of type bool that has explicit "empty" value.
type Bool struct {
	Value   bool
	Present bool
}

// MakeBool converts bool to Bool.
func MakeBool(value bool) Bool {
	return Bool{Value: value, Present: true}
}

// MakeBoolFromPtr converts a pointer to bool to Bool.
func MakeBoolFromPtr(ptr *bool) Bool {
	if ptr == nil {
		return Bool{}
	}

	return MakeBool(*ptr)
}

// SafeValue safely converts Bool to bool returning its default value if the value of Bool is not present.
func (o Bool) SafeValue() (value bool) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts Bool to a pointer to bool returning nil pointer if the value of Bool is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Bool.
func (o Bool) SafePtr() (ptr *bool) {
	if o.Present {
		ptr = new(bool)
		*ptr = o.Value
	}

	return
}

// SafeValueWithFlag safely converts Bool to bool returning it with a flag set to true
// or its default value and the flag set to false if the value of Bool is not present.
func (o Bool) SafeValueWithFlag() (value bool, ok bool) {
	if !o.Present {
		return
	}

	return o.Value, o.Present
}

// SafeValueWithErr safely converts Bool to bool returning it with nil error
// or its default value and non-nil error if the value of Bool is not present.
func (o Bool) SafeValueWithErr() (value bool, err error) {
	if !o.Present {
		err = errors.New("value of optional.Bool is not present")
		return
	}

	return o.Value, nil
}

// SafePtrWithFlag safely converts Bool to a pointer to bool returning a pointer to a value with a flag set to true
// or nil pointer and the flag set to false if the value of Bool is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Bool.
func (o Bool) SafePtrWithFlag() (ptr *bool, ok bool) {
	if !o.Present {
		return
	}

	ptr = new(bool)
	*ptr, ok = o.Value, o.Present

	return
}

// SafePtrWithErr safely converts Bool to a pointer to bool returning a pointer to a value with nil error
// or nil pointer and non-nil error if the value of Bool is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Bool.
func (o Bool) SafePtrWithErr() (ptr *bool, err error) {
	if !o.Present {
		err = errors.New("value of optional.Bool is not present")
		return
	}

	ptr = new(bool)
	*ptr = o.Value

	return
}

// Equals returns true if both values are not present or both values are present and are equal according to a provided determinant.
func (o Bool) Equals(other Bool, determinant func(this, other bool) bool) bool {
	return (!o.Present && !other.Present) || (o.Present && other.Present && determinant(o.Value, other.Value))
}

// Compare returns a result of provided comparator and true if both values are present, otherwise it returns 0 and false.
func (o Bool) Compare(other Bool, comparator func(this, other bool) int) (int, bool) {
	if !o.Present || !other.Present {
		return 0, false
	}

	return comparator(o.Value, other.Value), true
}

// Filter returns the Bool if its value is present and it matches the given predicate, otherwise it returns an empty Bool.
func (o Bool) Filter(test func(value bool) bool) (result Bool) {
	if !o.Present || !test(o.Value) {
		return
	}

	return o
}

// Map applies the provided mapping function to a value and returns its result as Bool if the value is present,
// otherwise is returns an empty Bool.
func (o Bool) Map(mapper func(value bool) (result bool, present bool)) (result Bool) {
	if !o.Present {
		return
	}

	result.Value, result.Present = mapper(o.Value)

	return
}

// IfPresent invokes the specified action with the value if it is present.
func (o Bool) IfPresent(action func(value bool)) {
	if o.Present {
		action(o.Value)
	}
}

// OrElse returns the value if it is present, otherwise it returns given other value.
func (o Bool) OrElse(other bool) bool {
	if !o.Present {
		return other
	}

	return o.Value
}

// OrElseWithFlag returns the value if it is present with a flag set to true, otherwise it returns given other value
// and the flag set to false.
func (o Bool) OrElseWithFlag(other bool) (bool, bool) {
	if !o.Present {
		return other, false
	}

	return o.Value, o.Present
}

// OrElseWithErr returns the value if it is present with nil error, otherwise it returns given other value
// and non-nil error.
func (o Bool) OrElseWithErr(other bool) (bool, error) {
	if !o.Present {
		return other, errors.New("value of optional.Bool is not present")
	}

	return o.Value, nil
}

// OrElseGet returns the value if it is present, otherwise it invokes a supplier and returns a result of that invocation.
func (o Bool) OrElseGet(supplier func() bool) bool {
	if !o.Present {
		return supplier()
	}

	return o.Value
}

// OrElseGetWithFlag returns the value if it is present with a flag set to true, otherwise it invokes a supplier and returns
// a result of that invocation with a flag set to false.
func (o Bool) OrElseGetWithFlag(supplier func() bool) (result bool, ok bool) {
	if !o.Present {
		return supplier(), false
	}

	return o.Value, o.Present
}

// OrElseGetWithErr returns the value if it is present with nil error, otherwise it invokes a supplier and returns
// a result of that invocation with non-nil error.
func (o Bool) OrElseGetWithErr(supplier func() bool) (result bool, err error) {
	if !o.Present {
		return supplier(), errors.New("value of optional.Bool is not present")
	}

	return o.Value, nil
}

// OrElseErr returns the value if it is present with nil error, otherwise it invokes an error supplier and returns
// default value and an error returned by the error supplier.
func (o Bool) OrElseErr(errSupplier func() error) (result bool, err error) {
	if !o.Present {
		err = errSupplier()
		return
	}

	return o.Value, nil
}

// MarshalJSON marshals Bool to json.
func (o Bool) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Bool from json.
func (o *Bool) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
