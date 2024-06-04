package models

import (
	"time"
)

type Event struct {
	PetID  string    `json:"petid"`
	Date   time.Time `json:"date" validate:"required"`
	Type   string    `json:"type" validate:"required"`
	Remark string    `json:"remark,omitempty"`
}
