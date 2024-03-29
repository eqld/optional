package optional

// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT

import (
	"encoding/json"
	"errors"
)

// Float32 is a superset of type float32 that has explicit "empty" value.
type Float32 struct {
	Value   float32
	Present bool
}

// MakeFloat32 converts float32 to Float32.
func MakeFloat32(value float32) Float32 {
	return Float32{Value: value, Present: true}
}

// MakeFloat32FromPtr converts a pointer to float32 to Float32.
func MakeFloat32FromPtr(ptr *float32) Float32 {
	if ptr == nil {
		return Float32{}
	}

	return MakeFloat32(*ptr)
}

// SafeValue safely converts Float32 to float32 returning its default value if the value of Float32 is not present.
func (o Float32) SafeValue() (value float32) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts Float32 to a pointer to float32 returning nil pointer if the value of Float32 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Float32.
func (o Float32) SafePtr() (ptr *float32) {
	if o.Present {
		ptr = new(float32)
		*ptr = o.Value
	}

	return
}

// SafeValueWithFlag safely converts Float32 to float32 returning it with a flag set to true
// or its default value and the flag set to false if the value of Float32 is not present.
func (o Float32) SafeValueWithFlag() (value float32, ok bool) {
	if !o.Present {
		return
	}

	return o.Value, o.Present
}

// SafeValueWithErr safely converts Float32 to float32 returning it with nil error
// or its default value and non-nil error if the value of Float32 is not present.
func (o Float32) SafeValueWithErr() (value float32, err error) {
	if !o.Present {
		err = errors.New("value of optional.Float32 is not present")
		return
	}

	return o.Value, nil
}

// SafePtrWithFlag safely converts Float32 to a pointer to float32 returning a pointer to a value with a flag set to true
// or nil pointer and the flag set to false if the value of Float32 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Float32.
func (o Float32) SafePtrWithFlag() (ptr *float32, ok bool) {
	if !o.Present {
		return
	}

	ptr = new(float32)
	*ptr, ok = o.Value, o.Present

	return
}

// SafePtrWithErr safely converts Float32 to a pointer to float32 returning a pointer to a value with nil error
// or nil pointer and non-nil error if the value of Float32 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Float32.
func (o Float32) SafePtrWithErr() (ptr *float32, err error) {
	if !o.Present {
		err = errors.New("value of optional.Float32 is not present")
		return
	}

	ptr = new(float32)
	*ptr = o.Value

	return
}

// Equals returns true if both values are not present or both values are present and are equal according to a provided determinant.
func (o Float32) Equals(other Float32, determinant func(this, other float32) bool) bool {
	return (!o.Present && !other.Present) || (o.Present && other.Present && determinant(o.Value, other.Value))
}

// Compare returns a result of provided comparator and true if both values are present, otherwise it returns 0 and false.
func (o Float32) Compare(other Float32, comparator func(this, other float32) int) (int, bool) {
	if !o.Present || !other.Present {
		return 0, false
	}

	return comparator(o.Value, other.Value), true
}

// Filter returns the Float32 if its value is present and it matches the given predicate, otherwise it returns an empty Float32.
func (o Float32) Filter(test func(value float32) bool) (result Float32) {
	if !o.Present || !test(o.Value) {
		return
	}

	return o
}

// Map applies the provided mapping function to a value and returns its result as Float32 if the value is present,
// otherwise is returns an empty Float32.
func (o Float32) Map(mapper func(value float32) (result float32, present bool)) (result Float32) {
	if !o.Present {
		return
	}

	result.Value, result.Present = mapper(o.Value)

	return
}

// IfPresent invokes the specified action with the value if it is present.
func (o Float32) IfPresent(action func(value float32)) {
	if o.Present {
		action(o.Value)
	}
}

// OrElse returns the value if it is present, otherwise it returns given other value.
func (o Float32) OrElse(other float32) float32 {
	if !o.Present {
		return other
	}

	return o.Value
}

// OrElseWithFlag returns the value if it is present with a flag set to true, otherwise it returns given other value
// and the flag set to false.
func (o Float32) OrElseWithFlag(other float32) (float32, bool) {
	if !o.Present {
		return other, false
	}

	return o.Value, o.Present
}

// OrElseWithErr returns the value if it is present with nil error, otherwise it returns given other value
// and non-nil error.
func (o Float32) OrElseWithErr(other float32) (float32, error) {
	if !o.Present {
		return other, errors.New("value of optional.Float32 is not present")
	}

	return o.Value, nil
}

// OrElseGet returns the value if it is present, otherwise it invokes a supplier and returns a result of that invocation.
func (o Float32) OrElseGet(supplier func() float32) float32 {
	if !o.Present {
		return supplier()
	}

	return o.Value
}

// OrElseGetWithFlag returns the value if it is present with a flag set to true, otherwise it invokes a supplier and returns
// a result of that invocation with a flag set to false.
func (o Float32) OrElseGetWithFlag(supplier func() float32) (result float32, ok bool) {
	if !o.Present {
		return supplier(), false
	}

	return o.Value, o.Present
}

// OrElseGetWithErr returns the value if it is present with nil error, otherwise it invokes a supplier and returns
// a result of that invocation with non-nil error.
func (o Float32) OrElseGetWithErr(supplier func() float32) (result float32, err error) {
	if !o.Present {
		return supplier(), errors.New("value of optional.Float32 is not present")
	}

	return o.Value, nil
}

// OrElseErr returns the value if it is present with nil error, otherwise it invokes an error supplier and returns
// default value and an error returned by the error supplier.
func (o Float32) OrElseErr(errSupplier func() error) (result float32, err error) {
	if !o.Present {
		err = errSupplier()
		return
	}

	return o.Value, nil
}

// MarshalJSON marshals Float32 to json.
func (o Float32) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Float32 from json.
func (o *Float32) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
