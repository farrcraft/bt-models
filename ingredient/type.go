package ingredient

// Type describes a type of ingredient
// -- e.g. grain, hop, yeast, herb, spice, fruit, fining, mineral, chemical, fermentable, adjunct, flavor
type Type struct {
	UUID string `json:"ingredient_type_uuid"`
	Name string `json:"name"`
}
