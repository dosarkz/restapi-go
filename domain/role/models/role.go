package models

const (
	roleAdmin      = "admin"
	roleManager    = "manager"
	RoleAdminInt   = 1
	RoleManagerInt = 2
)

type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type NewRole struct {
	Name string `json:"name"`
}

type UpdateRole struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
