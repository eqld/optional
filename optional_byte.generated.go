package optional

// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT

import (
	"encoding/json"
	"errors"
)

// Byte is a superset of type byte that has explicit "empty" value.
type Byte struct {
	Value   byte
	Present bool
}

// MakeByte converts byte to Byte.
func MakeByte(value byte) Byte {
	return Byte{Value: value, Present: true}
}

// MakeByteFromPtr converts a pointer to byte to Byte.
func MakeByteFromPtr(ptr *byte) Byte {
	if ptr == nil {
		return Byte{}
	}

	return MakeByte(*ptr)
}

// SafeValue safely converts Byte to byte returning its default value if the value of Byte is not present.
func (o Byte) SafeValue() (value byte) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts Byte to a pointer to byte returning nil pointer if the value of Byte is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Byte.
func (o Byte) SafePtr() (ptr *byte) {
	if o.Present {
		ptr = new(byte)
		*ptr = o.Value
	}

	return
}

// SafeValueWithFlag safely converts Byte to byte returning it with a flag set to true
// or its default value and the flag set to false if the value of Byte is not present.
func (o Byte) SafeValueWithFlag() (value byte, ok bool) {
	if !o.Present {
		return
	}

	return o.Value, o.Present
}

// SafeValueWithErr safely converts Byte to byte returning it with nil error
// or its default value and non-nil error if the value of Byte is not present.
func (o Byte) SafeValueWithErr() (value byte, err error) {
	if !o.Present {
		err = errors.New("value of optional.Byte is not present")
		return
	}

	return o.Value, nil
}

// SafePtrWithFlag safely converts Byte to a pointer to byte returning a pointer to a value with a flag set to true
// or nil pointer and the flag set to false if the value of Byte is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Byte.
func (o Byte) SafePtrWithFlag() (ptr *byte, ok bool) {
	if !o.Present {
		return
	}

	ptr = new(byte)
	*ptr, ok = o.Value, o.Present

	return
}

// SafePtrWithErr safely converts Byte to a pointer to byte returning a pointer to a value with nil error
// or nil pointer and non-nil error if the value of Byte is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Byte.
func (o Byte) SafePtrWithErr() (ptr *byte, err error) {
	if !o.Present {
		err = errors.New("value of optional.Byte is not present")
		return
	}

	ptr = new(byte)
	*ptr = o.Value

	return
}

// Equals returns true if both values are not present or both values are present and are equal according to a provided determinant.
func (o Byte) Equals(other Byte, determinant func(this, other byte) bool) bool {
	return (!o.Present && !other.Present) || (o.Present && other.Present && determinant(o.Value, other.Value))
}

// Compare returns a result of provided comparator and true if both values are present, otherwise it returns 0 and false.
func (o Byte) Compare(other Byte, comparator func(this, other byte) int) (int, bool) {
	if !o.Present || !other.Present {
		return 0, false
	}

	return comparator(o.Value, other.Value), true
}

// Filter returns the Byte if its value is present and it matches the given predicate, otherwise it returns an empty Byte.
func (o Byte) Filter(test func(value byte) bool) (result Byte) {
	if !o.Present || !test(o.Value) {
		return
	}

	return o
}

// Map applies the provided mapping function to a value and returns its result as Byte if the value is present,
// otherwise is returns an empty Byte.
func (o Byte) Map(mapper func(value byte) (result byte, present bool)) (result Byte) {
	if !o.Present {
		return
	}

	result.Value, result.Present = mapper(o.Value)

	return
}

// IfPresent invokes the specified action with the value if it is present.
func (o Byte) IfPresent(action func(value byte)) {
	if o.Present {
		action(o.Value)
	}
}

// OrElse returns the value if it is present, otherwise it returns given other value.
func (o Byte) OrElse(other byte) byte {
	if !o.Present {
		return other
	}

	return o.Value
}

// OrElseFlag returns the value if it is present with a flag set to true, otherwise it returns given other value
// and the flag set to false.
func (o Byte) OrElseFlag(other byte) (byte, bool) {
	if !o.Present {
		return other, false
	}

	return o.Value, o.Present
}

// OrElseErr returns the value if it is present with nil error, otherwise it returns given other value
// and non-nil error.
func (o Byte) OrElseErr(other byte) (byte, error) {
	if !o.Present {
		return other, errors.New("value of optional.Byte is not present")
	}

	return o.Value, nil
}

// OrElseGet returns the value if it is present, otherwise it invokes a supplier and returns a result of that invocation.
func (o Byte) OrElseGet(supplier func() byte) byte {
	if !o.Present {
		return supplier()
	}

	return o.Value
}

// OrElseGetFlag returns the value if it is present with a flag set to true, otherwise it invokes a supplier and returns
// a result of that invocation with a flag set to false.
func (o Byte) OrElseGetFlag(supplier func() byte) (result byte, ok bool) {
	if !o.Present {
		return supplier(), false
	}

	return o.Value, o.Present
}

// OrElseGetErr returns the value if it is present with nil error, otherwise it invokes a supplier and returns
// a result of that invocation with non-nil error.
func (o Byte) OrElseGetErr(supplier func() byte) (result byte, err error) {
	if !o.Present {
		return supplier(), errors.New("value of optional.Byte is not present")
	}

	return o.Value, nil
}

// MarshalJSON marshals Byte to json.
func (o Byte) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Byte from json.
func (o *Byte) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
