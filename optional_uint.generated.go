// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT
package optional

import "encoding/json"

// Uint is a superset of type uint that has explicit "empty" value.
type Uint struct {
	Value   uint
	Present bool
}

// MarshalJSON marshalls Uint to json
func (o Uint) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshalls Uint from json
func (o *Uint) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}
