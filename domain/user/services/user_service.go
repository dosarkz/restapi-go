package services

import (
	"rest/app/helpers/paginator"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	"rest/domain/user/repositories"
	"rest/domain/user/responses"
)

type IUserService interface {
	AddUser(input forms.NewUser) (*models.User, error)
	UpdateUser(id uint, input forms.UpdateUser) (*models.User, error)
	ShowUser(id uint) (*models.User, error)
	GetUsers(params map[string]string) ([]models.User, error)
	Paginate(params map[string]string) (*paginator.Paginator, error)
	DestroyUser(id uint) error
}

type UserService struct {
	userRepo repositories.UserRepository
}

func (u UserService) AddUser(input forms.NewUser) (*models.User, error) {
	user, err := u.userRepo.Create(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) UpdateUser(id uint, input forms.UpdateUser) (*models.User, error) {
	model, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	update, uErr := u.userRepo.Update(forms.UserBindingUpdate(&input, model))
	if uErr != nil {
		return nil, uErr
	}

	return update, nil
}

func (u UserService) ShowUser(id uint) (*models.User, error) {
	byId, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return byId, nil
}

func (u UserService) GetUsers(params map[string]string) ([]models.User, error) {
	us, _, err := u.userRepo.FindAll(params)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (u UserService) Paginate(params map[string]string) (*paginator.Paginator, error) {
	params["withPaginate"] = "true"
	us, pagination, err := u.userRepo.FindAll(params)
	if err != nil {
		return nil, err
	}

	if len(us) == 0 {
		return pagination, nil
	}

	pagination.Data = responses.BindingUserListResponse(us)
	return pagination, nil
}

func (u UserService) DestroyUser(id uint) error {
	model, err := u.userRepo.FindById(id)
	if err != nil {
		return err
	}

	delErr := u.userRepo.Delete(model.ID)
	if delErr != nil {
		return delErr
	}
	return nil
}

type ServiceConfig struct {
	UserRepo repositories.UserRepository
}

func NewUserService(s ServiceConfig) *UserService {
	return &UserService{
		s.UserRepo,
	}
}
