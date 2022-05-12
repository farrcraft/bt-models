package core

import (
	"bt-models/ingredient"
	"bt-models/fermentable"
)

type Ingredient struct {
	UUID               string              `json:"ingredient_uuid"`
	Type               *ingredient.Type     `json:"ingredient_type"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	FermentableProfile *fermentable.Profile `json:"fermentable_profile"`
	Hop                *Hop                `json:"hop"`
	Yeast              *Yeast              `json:"yeast"`
	Distributor        *Distributor        `json:"distributor"`
	User               *User               `json:"user"`
}
