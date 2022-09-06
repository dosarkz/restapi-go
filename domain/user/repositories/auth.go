package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"rest/app/helpers/auth"
	"rest/app/helpers/hash"
	"rest/database/redis"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	"rest/domain/user/responses"
	"net/http"
)

type AuthRepository interface {
	Login(form *forms.LoginForm, userRepo UserRepository) (*responses.UserResponse, error)
	Logout(req *http.Request) (bool, error)
	CreateAccessToken(forms.AccessToken) error
}

type AuthRepo struct {
	db *gorm.DB
}

func (r *AuthRepo) Login(form *forms.LoginForm, userRepo UserRepository) (*responses.UserResponse, error) {
	user, err := userRepo.FindBy(form.Email, "email")
	if err != nil {
		return nil, err
	}
	if user.StatusID != models.Activated {
		return nil, fmt.Errorf("your account not activated")
	}

	if hash.CheckPasswordHash(form.Password, user.Password) != true {
		return nil, fmt.Errorf("user credentials are failed")
	}
	token, errToken := auth.GenerateToken(user.ID)
	if errToken != nil {
		return nil, errToken
	}

	return forms.UserBindingResponse(user, token), err
}

func (r *AuthRepo) Logout(req *http.Request) (bool, error) {
	tokenString := req.Header.Get("Authorization")

	if tokenString == "" {
		return false, fmt.Errorf("auth tag not found")
	}
	rConf := new(redis.Config)
	redisConn := redis.ConnectToRedis(rConf)
	_, err := redisConn.SetValue(tokenString, tokenString, 0)
	if err != nil {
		return false, fmt.Errorf("logout token failed")
	}
	return true, nil
}

func (r *AuthRepo) CreateAccessToken(input forms.AccessToken) error {
	tx := r.db.Begin()
	if err := r.db.Create(&input).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db}
}
