package core

type Batch struct {
	UUID   string  `json:"batch_uuid"`
	Recipe *Recipe `json:"recipe"`
	Date   int64   `json:"date_brewed_timestamp"`
	Name   string  `json:"name"`
}
