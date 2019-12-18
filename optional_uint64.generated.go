package optional

// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT

import (
	"encoding/json"
	"errors"
)

// Uint64 is a superset of type uint64 that has explicit "empty" value.
type Uint64 struct {
	Value   uint64
	Present bool
}

// MakeUint64 converts uint64 to Uint64.
func MakeUint64(value uint64) Uint64 {
	return Uint64{Value: value, Present: true}
}

// MakeUint64FromPtr converts a pointer to uint64 to Uint64.
func MakeUint64FromPtr(ptr *uint64) Uint64 {
	if ptr == nil {
		return Uint64{}
	}

	return MakeUint64(*ptr)
}

// SafeValue safely converts Uint64 to uint64 returning its default value if the value of Uint64 is not present.
func (o Uint64) SafeValue() (value uint64) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts Uint64 to a pointer to uint64 returning nil pointer if the value of Uint64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Uint64.
func (o Uint64) SafePtr() (ptr *uint64) {
	if o.Present {
		ptr = new(uint64)
		*ptr = o.Value
	}

	return
}

// SafeValueWithFlag safely converts Uint64 to uint64 returning it with a flag set to true
// or its default value and the flag set to false if the value of Uint64 is not present.
func (o Uint64) SafeValueWithFlag() (value uint64, ok bool) {
	if !o.Present {
		return
	}

	return o.Value, o.Present
}

// SafeValueWithErr safely converts Uint64 to uint64 returning it with nil error
// or its default value and non-nil error if the value of Uint64 is not present.
func (o Uint64) SafeValueWithErr() (value uint64, err error) {
	if !o.Present {
		err = errors.New("value of optional.Uint64 is not present")
		return
	}

	return o.Value, nil
}

// SafePtrWithFlag safely converts Uint64 to a pointer to uint64 returning a pointer to a value with a flag set to true
// or nil pointer and the flag set to false if the value of Uint64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Uint64.
func (o Uint64) SafePtrWithFlag() (ptr *uint64, ok bool) {
	if !o.Present {
		return
	}

	ptr = new(uint64)
	*ptr, ok = o.Value, o.Present

	return
}

// SafePtrWithErr safely converts Uint64 to a pointer to uint64 returning a pointer to a value with nil error
// or nil pointer and non-nil error if the value of Uint64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Uint64.
func (o Uint64) SafePtrWithErr() (ptr *uint64, err error) {
	if !o.Present {
		err = errors.New("value of optional.Uint64 is not present")
		return
	}

	ptr = new(uint64)
	*ptr = o.Value

	return
}

// Equals returns true if both values are not present or both values are present and are equal according to a provided determinant.
func (o Uint64) Equals(other Uint64, determinant func(this, other uint64) bool) bool {
	return (!o.Present && !other.Present) || (o.Present && other.Present && determinant(o.Value, other.Value))
}

// Compare returns a result of provided comparator and true if both values are present, otherwise it returns 0 and false.
func (o Uint64) Compare(other Uint64, comparator func(this, other uint64) int) (int, bool) {
	if !o.Present || !other.Present {
		return 0, false
	}

	return comparator(o.Value, other.Value), true
}

// Filter returns the Uint64 if its value is present and it matches the given predicate, otherwise it returns an empty Uint64.
func (o Uint64) Filter(test func(value uint64) bool) (result Uint64) {
	if !o.Present || !test(o.Value) {
		return
	}

	return o
}

// Map applies the provided mapping function to a value and returns its result as Uint64 if the value is present,
// otherwise is returns an empty Uint64.
func (o Uint64) Map(mapper func(value uint64) (result uint64, present bool)) (result Uint64) {
	if !o.Present {
		return
	}

	result.Value, result.Present = mapper(o.Value)

	return
}

// IfPresent invokes the specified action with the value if it is present.
func (o Uint64) IfPresent(action func(value uint64)) {
	if o.Present {
		action(o.Value)
	}
}

// OrElse returns the value if it is present, otherwise it returns given other value.
func (o Uint64) OrElse(other uint64) uint64 {
	if !o.Present {
		return other
	}

	return o.Value
}

// OrElseWithFlag returns the value if it is present with a flag set to true, otherwise it returns given other value
// and the flag set to false.
func (o Uint64) OrElseWithFlag(other uint64) (uint64, bool) {
	if !o.Present {
		return other, false
	}

	return o.Value, o.Present
}

// OrElseWithErr returns the value if it is present with nil error, otherwise it returns given other value
// and non-nil error.
func (o Uint64) OrElseWithErr(other uint64) (uint64, error) {
	if !o.Present {
		return other, errors.New("value of optional.Uint64 is not present")
	}

	return o.Value, nil
}

// OrElseGet returns the value if it is present, otherwise it invokes a supplier and returns a result of that invocation.
func (o Uint64) OrElseGet(supplier func() uint64) uint64 {
	if !o.Present {
		return supplier()
	}

	return o.Value
}

// OrElseGetWithFlag returns the value if it is present with a flag set to true, otherwise it invokes a supplier and returns
// a result of that invocation with a flag set to false.
func (o Uint64) OrElseGetWithFlag(supplier func() uint64) (result uint64, ok bool) {
	if !o.Present {
		return supplier(), false
	}

	return o.Value, o.Present
}

// OrElseGetWithErr returns the value if it is present with nil error, otherwise it invokes a supplier and returns
// a result of that invocation with non-nil error.
func (o Uint64) OrElseGetWithErr(supplier func() uint64) (result uint64, err error) {
	if !o.Present {
		return supplier(), errors.New("value of optional.Uint64 is not present")
	}

	return o.Value, nil
}

// OrElseErr returns the value if it is present with nil error, otherwise it invokes an error supplier and returns
// default value and an error returned by the error supplier.
func (o Uint64) OrElseErr(errSupplier func() error) (result uint64, err error) {
	if !o.Present {
		err = errSupplier()
		return
	}

	return o.Value, nil
}

// MarshalJSON marshals Uint64 to json.
func (o Uint64) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Uint64 from json.
func (o *Uint64) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
