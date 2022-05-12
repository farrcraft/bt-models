package bjcp

import "bt-models/attribute"

// StyleAttributeValueMap describes attributes associated with a BJCP style
type StyleAttributeValueMap struct {
	UUID           string           `json:"bjcp_style_attribute_value_map"`
	BjcpStyle      *Style           `json:"bjcp_style"`
	AttributeValue *attribute.Value `json:"attribute_value"`
}
