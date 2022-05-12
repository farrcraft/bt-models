package core

// Vendor describes a generic vendor (name of any company that makes an ingredient)
// -- can also specifically be a maltster
type Vendor struct {
	UUID     string `json:"vendor_uuid"`
	Name     string `json:"name"`
	Maltster bool   `json:"maltster"`
}
