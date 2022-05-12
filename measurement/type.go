package measurement

// Type describes the purpose for a measurement
// Such as where in the process that the measurement was taken -- e.g. mash, boil, fermentation
type Type struct {
	UUID string `json:"measurement_type_uuid"`
	Name string `json:"name"`
}
