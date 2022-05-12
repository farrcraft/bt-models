package bjcp

// Style describes a BJCP style
type Style struct {
	UUID     string `json:"bjcp_style_uuid"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Parent   *Style `json:"parent_bjcp_style"`
}
