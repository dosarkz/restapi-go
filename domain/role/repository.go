package role

import (
	"gorm.io/gorm"
	"rest/domain/role/forms"
	"rest/domain/role/models"
)

type Repository interface {
	Create(input models.NewRole) (*models.Role, error)
	List() ([]*models.Role, error)
	Show(id uint) (*models.Role, error)
	Update(input models.UpdateRole) (*models.Role, error)
	Delete(id uint) error
}

type Repo struct {
	db *gorm.DB
}

func (r *Repo) Create(input models.NewRole) (*models.Role, error) {
	var role *models.Role
	role = forms.RoleBinding(&input)
	tx := r.db.Begin()
	if err := r.db.Create(&role).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return role, nil
}

func (r *Repo) List() ([]*models.Role, error) {
	var roles []*models.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *Repo) Show(id uint) (*models.Role, error) {
	var role models.Role
	if err := r.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Repo) Update(input models.UpdateRole) (*models.Role, error) {
	role, err := r.Show(input.ID)
	if err != nil {
		return nil, err
	}
	role = forms.RoleBindingUpdate(&input)
	tx := r.db.Begin()
	if err := r.db.Save(role).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return role, err
}

func (r *Repo) Delete(id uint) error {
	return r.db.Delete(&models.Role{}, id).Error
}

func NewRoleRepo(db *gorm.DB) *Repo {
	return &Repo{db}
}
