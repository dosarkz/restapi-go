package responses

import (
	role "rest/domain/role/models"
	"rest/domain/user/models"
	"time"
)

type UserModelResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone" gorm:"unique"`
	Role      role.Role `json:"role" gorm:"foreignKey:RoleID"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Email     string    `json:"email"`
}

type UserShortResponse struct {
	ID     uint      `json:"id" gorm:"primaryKey"`
	Name   string    `json:"name"`
	Phone  string    `json:"phone" gorm:"unique"`
	Role   role.Role `json:"role"`
	Status Status    `json:"status"`
	Email  string    `json:"email"`
}

func BindingUserShortModelResponse(entity models.User) *UserShortResponse {
	return &UserShortResponse{
		ID:     entity.ID,
		Name:   entity.Name,
		Phone:  entity.Phone,
		Role:   entity.Role,
		Status: *GetStatus(entity.StatusID),
		Email:  entity.Email,
	}
}

type Status struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func GetStatusNames() map[int]string {
	return map[int]string{
		models.Activated:   "Активный",
		models.Deactivated: "Не активный",
	}
}

func GetStatus(statusID int) *Status {
	return &Status{
		ID:    statusID,
		Title: GetStatusNames()[statusID],
	}
}

type UserResponse struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone" gorm:"unique"`
	Role      role.Role  `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Email     string     `json:"email"`
	Status    Status     `json:"status"`
	Token     string     `json:"token"`
}

func BindingUserModelResponse(entity models.User) *UserModelResponse {
	return &UserModelResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Phone:     entity.Phone,
		Role:      entity.Role,
		Email:     entity.Email,
		Status:    *GetStatus(entity.StatusID),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
