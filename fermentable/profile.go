package fermentable

import "bt-models/core"

// Profile describes the common properties of fermentable ingredients (grains, sugars, honeys, fruits, etc)
type Profile struct {
	UUID            string       `json:"fermentable_profile_uuid"`
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Origin          string       `json:"origin"`
	FermentableType *Type        `json:"fermentable_type"`
	Vendor          *core.Vendor `json:"vendor"`
}
