package unit

// Category describes the category that a unit of measurement can be grouped under
// e.g. temperature, weight, volume, time, specific gravity, pressure
type Category struct {
	UUID string `json:"unit_category_uuid"`
	Name string `json:"name"`
}
