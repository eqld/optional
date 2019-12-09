// Code generated by github.com/eqld/optional/cmd, DO NOT EDIT
package optional

import "encoding/json"

// Bool is a superset of type bool that has explicit "empty" value.
type Bool struct {
	Value   bool
	Present bool
}

// MarshalJSON marshalls Bool to json
func (o Bool) MarshalJSON() ([]byte, error) {
	if !o.Present {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value)
}

// UnmarshalJSON unmarshalls Bool from json
func (o *Bool) UnmarshalJSON(data []byte) error {
	if len(data) == 4 && string(data) == "null" {
		return nil
	}

	o.Present = true

	return json.Unmarshal(data, &o.Value)
}