package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"rest/app/helpers/paginator"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	"rest/domain/user/responses"
)

type UserRepository interface {
	Create(input forms.NewUser) (*responses.UserModelResponse, error)
	ListModel(params map[string]string) ([]*models.User, error)
	List(params map[string]string) ([]*responses.UserModelResponse, error)
	Show(id uint) (*models.User, error)
	Update(id uint, input forms.UpdateUser) (*models.User, error)
	UpdateModel(id uint, input forms.UpdateUser) (*responses.UserModelResponse, error)
	Delete(id uint) error
	FindBy(value string, field string) (*models.User, error)
	ShowById(id uint) (*models.User, error)
	ShowModel(id uint) (*responses.UserModelResponse, error)
	ShowByEmail(email string) (*models.User, error)
	Paginate(params map[string]string) (*paginator.Paginator, error)
}

type Repo struct {
	db *gorm.DB
}

func (r *Repo) Create(input forms.NewUser) (*responses.UserModelResponse, error) {
	userEntity := forms.UserBinding(&input)
	tx := r.db.Begin()
	if err := r.db.Create(userEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	userToReturn, err := r.ShowModel(userEntity.ID)
	if err != nil {
		return nil, err
	}
	return userToReturn, nil
}

func (r *Repo) ListModel(params map[string]string) ([]*models.User, error) {
	var users []*models.User
	filterQuery := models.Filter(r.db, params)
	sortQuery := models.Sort(filterQuery, params)
	if err := sortQuery.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repo) Paginate(params map[string]string) (*paginator.Paginator, error) {
	var pagination paginator.Paginator
	var users []*models.User
	dbQuery := r.db.Model(models.User{})
	dbQuery = models.Filter(r.db, params)
	dbQuery = models.Sort(dbQuery, params)
	if err := dbQuery.Scopes(paginator.PaginateFromMap(params, &pagination, dbQuery)).Find(&users).Error; err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return &pagination, nil
	}
	var usersToReturn []*responses.UserModelResponse
	for _, s := range users {
		usersToReturn = append(usersToReturn, forms.UserBindingModel(s))
	}
	pagination.Data = usersToReturn

	return &pagination, nil
}

func (r *Repo) List(params map[string]string) ([]*responses.UserModelResponse, error) {
	var users []*models.User
	filterQuery := models.Filter(r.db, params)
	sortQuery := models.Sort(filterQuery, params)
	if err := sortQuery.Find(&users).Error; err != nil {
		return nil, err
	}
	var usersToReturn []*responses.UserModelResponse
	for i := 0; i < len(users); i++ {
		userToReturn := forms.UserBindingModel(users[i])
		usersToReturn = append(usersToReturn, userToReturn)
	}
	return usersToReturn, nil
}

func (r *Repo) ShowModel(id uint) (*responses.UserModelResponse, error) {
	var user models.User
	if err := r.db.Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}
	userToReturn := forms.UserBindingModel(&user)
	return userToReturn, nil
}

func (r *Repo) Show(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).
		Preload("Role").
		Preload("Warehouses").
		Preload("Counterparties").
		Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) ShowById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) ShowByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Role").Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) Update(id uint, input forms.UpdateUser) (*models.User, error) {
	userToUpdate, err := r.ShowById(id)
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
	return userUpdated, err
}

func (r *Repo) UpdateModel(id uint, input forms.UpdateUser) (*responses.UserModelResponse, error) {
	userToUpdate, err := r.ShowById(id)
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
