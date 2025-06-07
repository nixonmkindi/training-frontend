package models

import "time"

// User defines class user data
type SubSystems struct {
	ID          int32     `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Code        string    `json:"code" form:"code"`
	IPAddress   string    `json:"ip_address" form:"ip_address"`
	Domain      string    `json:"domain" form:"domain"`
	Port        int32     `json:"port" form:"port"`
	Description string    `json:"description" form:"description"`
	Active      bool      `json:"active" form:"active"`
	CreatedBy   int32     `json:"created_by" form:"created_by"`
	UpdatedBy   int32     `json:"updated_by" form:"updated_by"`
	DeletedBy   int32     `json:"deleted_by" form:"deleted_by"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}
