package batch

import (
	"bt-models/core"
	"bt-models/equipment"
)

// Equipment describes the equipment used throughout the lifecycle of a batch
type Equipment struct {
	UUID             string             `json:"batch_equipment_uuid"`
	Batch            *core.Batch        `json:"batch"`
	EquipmentProfile *equipment.Profile `json:"equipment_profile"`
	EquipmentUsage   *equipment.Usage   `json:"equipment_usage"`
}
