package batch

import "bt-models/attribute"

// AttributeValueMap defines an attribute value for a batch
// These could be e.g. brewhouse efficiency, mash efficiency, actual og, actual fg
type AttributeValueMap struct {
	UUID           string           `json:"batch_attribute_value_map_uuid"`
	Batch          *Batch           `json:"batch"`
	AttributeValue *attribute.Value `json:"attribute_value"`
}
