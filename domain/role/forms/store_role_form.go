package forms

import (
	"rest/domain/role/models"
)

type StoreRoleForm struct {
	Name string
}

func RoleBinding(entity *models.NewRole) *models.Role {
	return &models.Role{
		Name: entity.Name,
	}
}

func RoleBindingUpdate(entity *models.UpdateRole) *models.Role {
	return &models.Role{
		Name: entity.Name,
	}
}
