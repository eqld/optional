package main

const template = `package _PKG_

// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT

import (
	"encoding/json"
	"errors"
)

// _TO_ is a superset of type _T_ that has explicit "empty" value.
type _TO_ struct {
	Value   _T_
	Present bool
}

// Make_TO_ converts _T_ to _TO_.
func Make_TO_(value _T_) _TO_ {
	return _TO_{Value: value, Present: true}
}

// Make_TO_FromPtr converts a pointer to _T_ to _TO_.
func Make_TO_FromPtr(ptr *_T_) _TO_ {
	if ptr == nil {
		return _TO_{}
	}

	return Make_TO_(*ptr)
}

// SafeValue safely converts _TO_ to _T_ returning its default value if the value of _TO_ is not present.
func (o _TO_) SafeValue() (value _T_) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts _TO_ to a pointer to _T_ returning nil pointer if the value of _TO_ is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of _TO_.
func (o _TO_) SafePtr() (ptr *_T_) {
	if o.Present {
		ptr = new(_T_)
		*ptr = o.Value
	}

	return
}

// SafeValueWithFlag safely converts _TO_ to _T_ returning it with a flag set to true
// or its default value and the flag set to false if the value of _TO_ is not present.
func (o _TO_) SafeValueWithFlag() (value _T_, ok bool) {
	if !o.Present {
		return
	}

	return o.Value, o.Present
}

// SafeValueWithErr safely converts _TO_ to _T_ returning it with nil error
// or its default value and non-nil error if the value of _TO_ is not present.
func (o _TO_) SafeValueWithErr() (value _T_, err error) {
	if !o.Present {
		err = errors.New("value of _PKG_._TO_ is not present")
		return
	}

	return o.Value, nil
}

// SafePtrWithFlag safely converts _TO_ to a pointer to _T_ returning a pointer to a value with a flag set to true
// or nil pointer and the flag set to false if the value of _TO_ is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of _TO_.
func (o _TO_) SafePtrWithFlag() (ptr *_T_, ok bool) {
	if !o.Present {
		return
	}

	ptr = new(_T_)
	*ptr, ok = o.Value, o.Present

	return
}

// SafePtrWithErr safely converts _TO_ to a pointer to _T_ returning a pointer to a value with nil error
// or nil pointer and non-nil error if the value of _TO_ is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of _TO_.
func (o _TO_) SafePtrWithErr() (ptr *_T_, err error) {
	if !o.Present {
		err = errors.New("value of _PKG_._TO_ is not present")
		return
	}

	ptr = new(_T_)
	*ptr = o.Value

	return
}

// Equals returns true if both values are not present or both values are present and are equal according to a provided determinant.
func (o _TO_) Equals(other _TO_, determinant func(this, other _T_) bool) bool {
	return (!o.Present && !other.Present) || (o.Present && other.Present && determinant(o.Value, other.Value))
}

// Compare returns a result of provided comparator and true if both values are present, otherwise it returns 0 and false.
func (o _TO_) Compare(other _TO_, comparator func(this, other _T_) int) (int, bool) {
	if !o.Present || !other.Present {
		return 0, false
	}

	return comparator(o.Value, other.Value), true
}

// Filter returns the _TO_ if its value is present and it matches the given predicate, otherwise it returns an empty _TO_.
func (o _TO_) Filter(test func(value _T_) bool) (result _TO_) {
	if !o.Present || !test(o.Value) {
		return
	}

	return o
}

// Map applies the provided mapping function to a value and returns its result as _TO_ if the value is present,
// otherwise is returns an empty _TO_.
func (o _TO_) Map(mapper func(value _T_) (result _T_, present bool)) (result _TO_) {
	if !o.Present {
		return
	}

	result.Value, result.Present = mapper(o.Value)

	return
}

// IfPresent invokes the specified action with the value if it is present.
func (o _TO_) IfPresent(action func(value _T_)) {
	if o.Present {
		action(o.Value)
	}
}

// OrElse returns the value if it is present, otherwise it returns given other value.
func (o _TO_) OrElse(other _T_) _T_ {
	if !o.Present {
		return other
	}

	return o.Value
}

// OrElseWithFlag returns the value if it is present with a flag set to true, otherwise it returns given other value
// and the flag set to false.
func (o _TO_) OrElseWithFlag(other _T_) (_T_, bool) {
	if !o.Present {
		return other, false
	}

	return o.Value, o.Present
}

// OrElseWithErr returns the value if it is present with nil error, otherwise it returns given other value
// and non-nil error.
func (o _TO_) OrElseWithErr(other _T_) (_T_, error) {
	if !o.Present {
		return other, errors.New("value of _PKG_._TO_ is not present")
	}

	return o.Value, nil
}

// OrElseGet returns the value if it is present, otherwise it invokes a supplier and returns a result of that invocation.
func (o _TO_) OrElseGet(supplier func() _T_) _T_ {
	if !o.Present {
		return supplier()
	}

	return o.Value
}

// OrElseGetWithFlag returns the value if it is present with a flag set to true, otherwise it invokes a supplier and returns
// a result of that invocation with a flag set to false.
func (o _TO_) OrElseGetWithFlag(supplier func() _T_) (result _T_, ok bool) {
	if !o.Present {
		return supplier(), false
	}

	return o.Value, o.Present
}

// OrElseGetWithErr returns the value if it is present with nil error, otherwise it invokes a supplier and returns
// a result of that invocation with non-nil error.
func (o _TO_) OrElseGetWithErr(supplier func() _T_) (result _T_, err error) {
	if !o.Present {
		return supplier(), errors.New("value of _PKG_._TO_ is not present")
	}

	return o.Value, nil
}

// OrElseErr returns the value if it is present with nil error, otherwise it invokes an error supplier and returns
// default value and an error returned by the error supplier.
func (o _TO_) OrElseErr(errSupplier func() error) (result _T_, err error) {
	if !o.Present {
		err = errSupplier()
		return
	}

	return o.Value, nil
}

// MarshalJSON marshals _TO_ to json.
func (o _TO_) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals _TO_ from json.
func (o *_TO_) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
`
