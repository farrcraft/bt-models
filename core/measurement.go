package core

import (
	"bt-models/attribute"
	"bt-models/measurement"
)

type Measurement struct {
	UUID            string            `json:"measurement_uuid"`
	AttributeValue  *attribute.Value  `json:"attribute_value"`
	MeasurementType *measurement.Type `json:"measurement_type"`
	User            *User             `json:"user"`
	Batch           *Batch            `json:"batch"`
	Date            int64             `json:"date_measured_timestamp"`
}
