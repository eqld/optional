// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT
package optional

import "encoding/json"

// Uint64 is a superset of type uint64 that has explicit "empty" value.
type Uint64 struct {
	Value   uint64
	Present bool
}

// MakeUint64 converts uint64 to Uint64
func MakeUint64(value uint64) Uint64 {
	return Uint64{Value: value, Present: true}
}

// MakeUint64FromPtr converts a pointer to uint64 to Uint64
func MakeUint64FromPtr(ptr *uint64) Uint64 {
	if ptr == nil {
		return Uint64{}
	}

	return MakeUint64(*ptr)
}

// SafeValue safely converts Uint64 to uint64 returning its default value if the value of Uint64 is not present
func (o Uint64) SafeValue() (value uint64) {
	if o.Present {
		value = o.Value
	}

	return
}

// SafePtr safely converts Uint64 to a pointer to uint64 returning nil pointer if the value of Uint64 is not present;
// the pointer, if not nil, DOES NOT point to the underlying value of Uint64
func (o Uint64) SafePtr() (ptr *uint64) {
	if o.Present {
		ptr = new(uint64)
		*ptr = o.Value
	}

	return
}

// MarshalJSON marshals Uint64 to json
func (o Uint64) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Uint64 from json
func (o *Uint64) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
