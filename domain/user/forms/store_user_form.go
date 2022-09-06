package forms

import (
	"rest/domain/user/models"
	"rest/domain/user/responses"
	"time"
)

func UserBindingResponse(entity *models.User, token string) *responses.UserResponse {
	return &responses.UserResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Phone:     entity.Phone,
		Role:      entity.Role,
		Status:    *responses.GetStatus(entity.StatusID),
		Token:     token,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func UserBindingUpdate(input *UpdateUser, user *models.User) *models.User {
	updatedAt := time.Now()
	return &models.User{
		ID:        user.ID,
		Name:      input.Name,
		Phone:     input.Phone,
		RoleID:    user.RoleID,
		Role:      user.Role,
		Password:  input.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
		Email:     input.Email,
		StatusID:  models.Activated,
	}
}

func UserBindingModel(input *models.User) *responses.UserModelResponse {
	updatedAt := time.Now()
	return &responses.UserModelResponse{
		ID:        input.ID,
		Name:      input.Name,
		Phone:     input.Phone,
		Role:      input.Role,
		Status:    *responses.GetStatus(input.StatusID),
		CreatedAt: input.CreatedAt,
		UpdatedAt: updatedAt,
		Email:     input.Email,
	}
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"required,gte=6"`
	RoleID   uint   `json:"roleID" validate:"omitempty,exists:roles-id"`
	Email    string `json:"email" validate:"required,doesnt_exist=users-email"`
	StatusID int    `json:"statusId"`
}

type UpdateUser struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"gte=6"`
	Email    string `json:"email" validate:"required,doesnt_exist=users-email"`
	StatusID int    `json:"statusId"`
}

type LoginForm struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gte=6"`
}
