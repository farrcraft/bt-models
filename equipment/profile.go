package equipment

// Profile describes a piece of equipment
type Profile struct {
	UUID          string `json:"equipment_profile_uuid"`
	User          *User  `json:"user"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	EquipmentType *Type  `json:"equipment_type"`
}
