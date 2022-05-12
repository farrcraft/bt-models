package unit

// Conversion describes how to convert between two unit kinds
type Conversion struct {
	UUID           string `json:"unit_conversion_uuid"`
	SourceUnitKind *Kind  `json:"source_unit_kind"`
	TargetUnitKind *Kind  `json:"target_unit_kind"`
	Factor         string `json:"factor"`
}
