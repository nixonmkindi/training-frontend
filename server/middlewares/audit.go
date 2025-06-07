package middlewares

import (
	"time"
)

type AuditTrails struct {
	ID        int32     `json:"id,omitempty"`
	Email     string    `json:"email"`
	IPAddress string    `json:"ip_address"`
	Client    string    `json:"client"`
	UserID    int32     `json:"user_id"`
	Action    string    `json:"action"`
	Method    string    `json:"method"`
	Url       string    `json:"url"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	DeletedBy int32     `json:"deleted_by"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
