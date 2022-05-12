package fermentable

import "bt-models/attribute"

// AttributeValueMap describes attributes of fermentables
// -- color (lovibond or srm), yield (percent dry yield or raw yield by weight), moisture (%),
// -- diastatic power (lintner), coarse/fine yield delta (grains & adjuncts), protein (%),
// -- post-boil (bool), recommend mash (bool), max in batch (% by weight)
type AttributeValueMap struct {
	UUID string `json:"fermentable_attribute_value_map_uuid"`
	FermentableProfile * Profile `json:"fermentable_profile"`
	AttributeValue * attribute.Value `json:"attribute_value"`
}
