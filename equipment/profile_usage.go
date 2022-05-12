package equipment

// ProfileUsage defines how different pieces of equipment can be used
// -- single pieces of equipment might have multiple potential usages
// -- e.g. kegs can be used for both fermentation and serving
type ProfileUsage struct {
	UUID             string   `json:"equipment_profile_usage_uuid"`
	EquipmentProfile *Profile `json:"equipment_profile"`
	EquipmentUsage   *Usage   `json:"equipment_usage"`
}
