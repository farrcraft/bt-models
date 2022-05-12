package unit

import "bt-models/measurement"

// Kind describes a kind of unit
type Kind struct {
	UUID              string              `json:"unit_kind_uuid"`
	Abbreviation      string              `json:"abbreviation"`
	Name              string              `json:"name"`
	UnitCategory      *Category           `json:"unit_category"`
	MeasurementSystem *measurement.System `json:"measurement_system"`
	State             string              `json:"state"` // liquid, solid
}
