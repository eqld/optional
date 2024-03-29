package optional

// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT

import (
	"encoding/json"
	"errors"
)

// Uintptr is a superset of type uintptr that has explicit "empty" value.
type Uintptr struct {
	Value   uintptr
	Present bool
}

// MakeUintptr converts uintptr to Uintptr.
func MakeUintptr(value uintptr) Uintptr {
	return Uintptr{Value: value, Present: true}
}

// MakeUintptrFromPtr converts a pointer to uintptr to Uintptr.
func MakeUintptrFromPtr(ptr *uintptr) Uintptr {
	if ptr == nil {
		return Uintptr{}
	}

	return MakeUintptr(*ptr)
}

// SafeValue safely converts Uintptr to uintptr returning its default value if the value of Uintptr is not present.
func (o Uintptr) SafeValue() (value uintptr) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts Uintptr to a pointer to uintptr returning nil pointer if the value of Uintptr is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Uintptr.
func (o Uintptr) SafePtr() (ptr *uintptr) {
	if o.Present {
		ptr = new(uintptr)
		*ptr = o.Value
	}

	return
}

// SafeValueWithFlag safely converts Uintptr to uintptr returning it with a flag set to true
// or its default value and the flag set to false if the value of Uintptr is not present.
func (o Uintptr) SafeValueWithFlag() (value uintptr, ok bool) {
	if !o.Present {
		return
	}

	return o.Value, o.Present
}

// SafeValueWithErr safely converts Uintptr to uintptr returning it with nil error
// or its default value and non-nil error if the value of Uintptr is not present.
func (o Uintptr) SafeValueWithErr() (value uintptr, err error) {
	if !o.Present {
		err = errors.New("value of optional.Uintptr is not present")
		return
	}

	return o.Value, nil
}

// SafePtrWithFlag safely converts Uintptr to a pointer to uintptr returning a pointer to a value with a flag set to true
// or nil pointer and the flag set to false if the value of Uintptr is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Uintptr.
func (o Uintptr) SafePtrWithFlag() (ptr *uintptr, ok bool) {
	if !o.Present {
		return
	}

	ptr = new(uintptr)
	*ptr, ok = o.Value, o.Present

	return
}

// SafePtrWithErr safely converts Uintptr to a pointer to uintptr returning a pointer to a value with nil error
// or nil pointer and non-nil error if the value of Uintptr is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Uintptr.
func (o Uintptr) SafePtrWithErr() (ptr *uintptr, err error) {
	if !o.Present {
		err = errors.New("value of optional.Uintptr is not present")
		return
	}

	ptr = new(uintptr)
	*ptr = o.Value

	return
}

// Equals returns true if both values are not present or both values are present and are equal according to a provided determinant.
func (o Uintptr) Equals(other Uintptr, determinant func(this, other uintptr) bool) bool {
	return (!o.Present && !other.Present) || (o.Present && other.Present && determinant(o.Value, other.Value))
}

// Compare returns a result of provided comparator and true if both values are present, otherwise it returns 0 and false.
func (o Uintptr) Compare(other Uintptr, comparator func(this, other uintptr) int) (int, bool) {
	if !o.Present || !other.Present {
		return 0, false
	}

	return comparator(o.Value, other.Value), true
}

// Filter returns the Uintptr if its value is present and it matches the given predicate, otherwise it returns an empty Uintptr.
func (o Uintptr) Filter(test func(value uintptr) bool) (result Uintptr) {
	if !o.Present || !test(o.Value) {
		return
	}

	return o
}

// Map applies the provided mapping function to a value and returns its result as Uintptr if the value is present,
// otherwise is returns an empty Uintptr.
func (o Uintptr) Map(mapper func(value uintptr) (result uintptr, present bool)) (result Uintptr) {
	if !o.Present {
		return
	}

	result.Value, result.Present = mapper(o.Value)

	return
}

// IfPresent invokes the specified action with the value if it is present.
func (o Uintptr) IfPresent(action func(value uintptr)) {
	if o.Present {
		action(o.Value)
	}
}

// OrElse returns the value if it is present, otherwise it returns given other value.
func (o Uintptr) OrElse(other uintptr) uintptr {
	if !o.Present {
		return other
	}

	return o.Value
}

// OrElseWithFlag returns the value if it is present with a flag set to true, otherwise it returns given other value
// and the flag set to false.
func (o Uintptr) OrElseWithFlag(other uintptr) (uintptr, bool) {
	if !o.Present {
		return other, false
	}

	return o.Value, o.Present
}

// OrElseWithErr returns the value if it is present with nil error, otherwise it returns given other value
// and non-nil error.
func (o Uintptr) OrElseWithErr(other uintptr) (uintptr, error) {
	if !o.Present {
		return other, errors.New("value of optional.Uintptr is not present")
	}

	return o.Value, nil
}

// OrElseGet returns the value if it is present, otherwise it invokes a supplier and returns a result of that invocation.
func (o Uintptr) OrElseGet(supplier func() uintptr) uintptr {
	if !o.Present {
		return supplier()
	}

	return o.Value
}

// OrElseGetWithFlag returns the value if it is present with a flag set to true, otherwise it invokes a supplier and returns
// a result of that invocation with a flag set to false.
func (o Uintptr) OrElseGetWithFlag(supplier func() uintptr) (result uintptr, ok bool) {
	if !o.Present {
		return supplier(), false
	}

	return o.Value, o.Present
}

// OrElseGetWithErr returns the value if it is present with nil error, otherwise it invokes a supplier and returns
// a result of that invocation with non-nil error.
func (o Uintptr) OrElseGetWithErr(supplier func() uintptr) (result uintptr, err error) {
	if !o.Present {
		return supplier(), errors.New("value of optional.Uintptr is not present")
	}

	return o.Value, nil
}

// OrElseErr returns the value if it is present with nil error, otherwise it invokes an error supplier and returns
// default value and an error returned by the error supplier.
func (o Uintptr) OrElseErr(errSupplier func() error) (result uintptr, err error) {
	if !o.Present {
		err = errSupplier()
		return
	}

	return o.Value, nil
}

// MarshalJSON marshals Uintptr to json.
func (o Uintptr) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Uintptr from json.
func (o *Uintptr) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
