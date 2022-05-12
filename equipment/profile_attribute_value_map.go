package equipment

import "bt-models/attribute"

// ProfileAttributeValueMap describes an attribute associated with a piece of equipment
// -- equipment attribute value kinds:
// -- equipment capacity (kettle volume / mash tun volume / etc)
// -- mash tun weight, mash tun specific heat, top up water volume, 
// -- equipment material (stainless/glass/plastic/etc)
// -- trub / chiller loss (volume lost from boiler after transfer to fermentation vessel)
// -- boil evaporation rate, boil time, lauter deadspace, kettle top up volume, hop utilization
type ProfileAttributeValueMap struct {
	UUID string `json:"equipment_profile_attribute_value_map_uuid"`
	EquipmentProfile * Profile `json:"equipment_profile"`
	AttributeValue * attribute.Value `json:"attribute_value"`
}
