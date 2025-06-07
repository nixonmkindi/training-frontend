package models

type UserRolePermission struct {
	User       *User         `json:"user"`
	Role       []*Role       `json:"role"`
	Permission []*Permission `json:"permission"`
}

type RoleUserPermission struct {
	Role       *Role         `json:"role"`
	User       []*User       `json:"user"`
	Permission []*Permission `json:"permission"`
}
