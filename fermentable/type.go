package fermentable

// Type describes types of fermentables
// -- grain, sugar, (liquid) extract, dry extract, adjunct
type Type struct {
	UUID string `json:"fermentable_type_uuid"`
	Name string `json:"name"`
}
