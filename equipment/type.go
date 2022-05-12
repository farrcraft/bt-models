package equipment

// Type describes equipment types
// e.g. -- keg, kettle, bucket, plastic bottle, round cooler, carboy, rectangular cooler, keggle, conical
type Type struct {
	UUID string `json:"equipment_type_uuid"`
	Name string `json:"name"`
}
