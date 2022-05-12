package boil

// Step describes an individual step as a part of a boil schedule
type Step struct {
	UUID string `json:"boil_step_uuid"`
	Schedule *Schedule `json:"boil_schedule"`
	StepTime int64 `json:"step_time"`
	Name string `json:"name"`
	Description string `json:"description"`
	Ingredient * Ingredient `json:"ingredient"`
	IngredientAmount float64 `json:"ingredient_amount"`
}
