package boil

// Schedule describes the amount of time for a boil step
type Schedule struct {
	UUID     string  `json:"boil_schedule_uuid"`
	Recipe   *Recipe `json:"recipe"`
	Duration int64   `json:"duration"`
}
