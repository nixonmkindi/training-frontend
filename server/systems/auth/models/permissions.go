package models

import "time"

type Permission struct {
	ID          string    `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Path        string    `json:"path" form:"path" validate:"required"`
	SubSystemID int32     `json:"sub_system_id" form:"sub_system_id"`
	Method      string    `json:"method" form:"method" validate:"required"`
	Service     string    `json:"service" form:"service" validate:"required"`
	SubService  string    `json:"sub_service" form:"sub_service" validate:"required"`
	Action      string    `json:"action" form:"action" validate:"required"`
	CreatedBy   int32     `json:"created_by" form:"created_by"`
	UpdatedBy   int32     `json:"updated_by" form:"updated_by"`
	DeletedBy   int32     `json:"deleted_by" form:"deleted_by"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}
