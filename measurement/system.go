package measurement

// System describes the system that a measurement is taken in (e.g. imperial, metric)
type System struct {
	UUID string `json:"measurement_system_uuid"`
	Name string `json:"name"`
}
