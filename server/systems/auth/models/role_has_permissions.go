package models

type HasPermission struct {
	ID           int32    `json:"id" form:"id"`
	PermissionID []string `json:"permission" form:"permission"`
}
