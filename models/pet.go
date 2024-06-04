package models

import (
	"time"
)

type Pet struct {
	ID      int        `json:"id,omitempty"`
	Name    string     `json:"name" validate:"required"`
	Owner   string     `json:"owner" validate:"required"`
	Species string     `json:"species" validate:"required"`
	Birth   *time.Time `json:"birth,omitempty"`
	Death   *time.Time `json:"death,omitempty"`
}
