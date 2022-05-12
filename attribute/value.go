package attribute

import "bt-models/unit"

// Value represents the actual value of an attribute instance
type Value struct {
	UUID         string     `json:"attribute_value_uuid"`
	Kind         *Kind      `json:"attribute_kind"`
	TextValue    string     `json:"t_value"`
	IntegerValue int64      `json:"n_value"`
	BooleanValue bool       `json:"b_value"`
	FloatValue   float64    `json:"f_value"`
	UnitKind     *unit.Kind `json:"unit_kind"`
	RangeType    string     `json:"range_type"` // minimum, maximum
}
