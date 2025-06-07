package models

import (
	"time"
)

//ApplicationLog DataStructure
type ApplicationLog struct {
	ID        string    `json:"id,omitempty" form:"id"`
	Level     string    `json:"level" form:"level"`
	Host      string    `json:"host" form:"host"`
	Process   string    `json:"process" form:"process"`
	Message   string    `json:"message" form:"message"`
	Exception string    `json:"exception" form:"exception"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at_at,omitempty"`
}
 