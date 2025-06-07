package models

import (
	"time"
)

type Login struct {
	ID       int32  `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password"  validate:"required,min=1"`
}

// Defines UserStructure

type UserRole struct {
	ID    int32   `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Roles []int32 `json:"roles" form:"roles[]" validate:"required"`
}

// User defines class user data
type User struct {
	ID              int32     `json:"id" form:"id"`
	UserCategoryID  int32     `json:"user_category_id" form:"user_category_id"`
	TelegramHandle  string    `json:"telegram_handle" form:"telegram_handle"`
	IsActive        bool      `json:"is_active" form:"is_active"`
	IsPasswordReset bool      `json:"is_password_reset" form:"is_password_reset"`
	Token           string    `json:"token" form:"token"`
	Name            string    `json:"name" form:"name"`
	Email           string    `json:"email" form:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at" form:"email_verified_at"`
	Password        string    `json:"password" form:"password"`
	LoginAttempt    int32     `json:"login_attempt" form:"login_attempt"`
	RememberToken   string    `json:"remember_token" form:"remember_token"`
	CampusID        int32     `json:"campus_id" form:"campus_id"`
	UniqueID        string    `json:"unique_id" form:"unique_id"`
	IsHOD           bool      `json:"is_hod" form:"is_hod"`
	CreatedBy       int32     `json:"created_by" form:"created_by"`
	UpdatedBy       int32     `json:"updated_by" form:"updated_by"`
	DeletedBy       int32     `json:"deleted_by" form:"deleted_by"`
	CreatedAt       time.Time `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at" form:"deleted_at"`
}
