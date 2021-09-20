package model

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRoles struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}
