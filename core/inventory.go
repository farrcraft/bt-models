package core

type Inventory struct {
	UUID string `json:"inventory_uuid"`
	User *User  `json:"user"`
}
