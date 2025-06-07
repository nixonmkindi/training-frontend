package models

type UserHasRole struct {
	RoleID []int32 `json:"role_id" form:"role_id"`
	UserID int32 `json:"user_id" form:"user_id"`
}
