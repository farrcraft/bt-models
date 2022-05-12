package ingredient

// Usage describes how an ingredient can be used
// -- boil, mash, primary, secondary, bottling
type Usage struct {
	UUID string `json:"ingredient_usage_uuid"`
	Name string `json:"name"`
}
