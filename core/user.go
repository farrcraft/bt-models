package core

import "bt-models/measurement"

type User struct {
	UUID              string              `json:"user_uuid"`
	Email             string              `json:"email"`
	FirstName         string              `json:"first_name"`
	LastName          string              `json:"last_name"`
	Created           int64               `json:"date_created_timestamp"`
	Updated           int64               `json:"date_updated_timestamp"`
	Deleted           bool                `json:"deleted"`
	MeasurementSystem *measurement.System `json:"measurement_system"`
	Password          string
	Salt              string
}
