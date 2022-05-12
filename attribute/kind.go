package attribute

// Kind describes a kind of attribute
type Kind struct {
	UUID         string `json:"attribute_kind_uuid"`
	Abbreviation string `json:"abbreviaton"`
	Description  string `json:"description"`
	ValueType    string `json:"value_type"` // string, integer, boolean, float
}
