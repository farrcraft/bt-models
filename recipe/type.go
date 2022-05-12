package recipe

// Type of recipe
// -- extract, partial mash, all grain
type Type struct {
	UUID string `json:"recipe_type_uuid"`
	Name string `json:"name"`
}
