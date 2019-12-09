// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT
package optional

import "encoding/json"

// Byte is a superset of type byte that has explicit "empty" value.
type Byte struct {
	Value   byte
	Present bool
}

// MarshalJSON marshalls Byte to json
func (o Byte) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshalls Byte from json
func (o *Byte) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}