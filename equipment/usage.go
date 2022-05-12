package equipment

// Usage describes how equipment can be used
// -- boil, mash, fermentation, serving, lauter
type Usage struct {
	UUID string `json:"equipment_usage_uuid"`
	Name string `json:"name"`
}
