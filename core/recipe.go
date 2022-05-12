package core

import (
	"bt-models/bjcp"
	"bt-models/recipe"
)

// Recipe describes a recipe
type Recipe struct {
	UUID          string       `json:"recipe_uuid"`
	User          *User        `json:"user"`
	Name          string       `json:"name"`
	RecipeType    *recipe.Type `json:"recipe_type"`
	LineageRecipe *Recipe      `json:"lineage_recipe"`
	BjcpStyle     *bjcp.Style  `json:"bjcp_style"`
	Date          int64        `json:"date_created_timestamp"`
	Visibility    string       `json:"visibility"` // public, private
}
