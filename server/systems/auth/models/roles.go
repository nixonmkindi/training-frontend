package models

import (
	"time"
)

// Role DataStructure
type Role struct {
	ID          int32     `json:"id,omitempty" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	CreatedBy   int32     `json:"created_by" form:"created_by"`
	UpdatedBy   int32     `json:"updated_by" form:"updated_by"`
	DeletedBy   int32     `json:"deleted_by" form:"deleted_by"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}
type RolePermissions struct {
	ID          string   `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Permissions []string `json:"permissions" form:"permissions[]" validate:"required"`
}
