package models

import (
	"time"
)

type Position struct {
	ID          int32     `json:"id" form:"id" validate:"omitempty,numeric"`
	Name        string    `json:"name,omitempty" form:"name" validate:"required"`
	Description string    `json:"description,omitempty" form:"description" validate:"omitempty"`
	CreatedBy   int32     `json:"created_by,omitempty" form:"created_by" validate:"omitempty,numeric"`
	UpdatedBy   int32     `json:"updated_by,omitempty" form:"updated_by" validate:"omitempty,numeric"`
	DeletedBy   int32     `json:"deleted_by,omitempty" form:"deleted_by" validate:"omitempty,numeric"`
	CreatedAt   time.Time `json:"created_at,omitempty" validate:"omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" validate:"omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty" validate:"omitempty"`
}
