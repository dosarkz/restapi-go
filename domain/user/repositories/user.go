package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"rest/app/helpers/paginator"
	roleModel "rest/domain/role/models"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	"rest/domain/user/responses"
)

type UserRepository interface {
	Create(input forms.NewUser) (*models.User, error)
	FindAll(params map[string]string) ([]models.User, *paginator.Paginator, error)
	FindById(id uint) (*models.User, error)
	Update(model *models.User) (*models.User, error)
	Delete(id uint) error
	FindBy(value string, field string) (*models.User, error)
}

type Repo struct {
	db *gorm.DB
}

func (r *Repo) Create(input forms.NewUser) (*models.User, error) {
	userEntity := models.User{
		RoleID: roleModel.RoleManagerInt,
	}
	tx := r.db.Begin()
	if err := r.db.Model(&userEntity).Preload("Role").Create(input).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &userEntity, nil
}

func (r *Repo) FindAll(params map[string]string) ([]models.User, *paginator.Paginator, error) {
	var (
		pagination paginator.Paginator
		users      []models.User
	)

	dbQuery := r.db.Model(models.User{}).Where("deleted_at IS NULL").Preload("Role")
	dbQuery = models.Filter(dbQuery, params)
	dbQuery = models.Sort(dbQuery, params)

	if err := dbQuery.Scopes(paginator.PaginateFromMap(params, &pagination, dbQuery)).
		Find(&users).Debug().Error; err != nil {
		return nil, nil, err
	}

	if len(users) == 0 {
		return nil, &pagination, nil
	}

	return users, &pagination, nil
}

func (r *Repo) FindById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).
		Preload("Role").
		Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) Update(model *models.User) (*models.User, error) {
	tx := r.db.Begin()
	if err := r.db.Preload("Role").Where("id = ? AND deleted_at IS NULL", model.ID).Save(&model).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return model, nil
}

func (r *Repo) UpdateModel(id uint, input forms.UpdateUser) (*responses.UserModelResponse, error) {
	userToUpdate, err := r.FindById(id)
	if err != nil {
		return nil, err
	}

	userUpdated := forms.UserBindingUpdate(&input, userToUpdate)

	tx := r.db.Begin()
	if err := r.db.Preload("Role").Where("id = ? AND deleted_at IS NULL", id).Save(&userUpdated).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	userToReturn := forms.UserBindingModel(userUpdated)
	return userToReturn, err
}

func (r *Repo) Delete(id uint) error {
	var user models.User
	tx := r.db.Begin()
	if err := r.db.Model(&user).Where("id", id).Update("deleted_at", "now()").Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *Repo) FindBy(value string, field string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").First(&user, fmt.Sprintf("%s = ?", field), value).Error
	return &user, err
}

func NewUserRepo(db *gorm.DB) *Repo {
	return &Repo{db}
}
