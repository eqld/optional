// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT
package optional

import "encoding/json"

// Uint8 is a superset of type uint8 that has explicit "empty" value.
type Uint8 struct {
	Value   uint8
	Present bool
}

// MakeUint8 converts uint8 to Uint8
func MakeUint8(value uint8) Uint8 {
	return Uint8{Value: value, Present: true}
}

// MarshalJSON marshals Uint8 to json
func (o Uint8) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}

	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshals Uint8 from json
func (o *Uint8) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
