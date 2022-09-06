package email

import (
	"rest/app/helpers/auth"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	"time"
)

const (
	Register      = 0
	ResetPassword = 1
)

func CreateAccessToken(user *models.User, tokenType uint) (*forms.AccessToken, error) {
	var accessToken forms.AccessToken
	var err error
	accessToken.UserID = user.ID
	accessToken.TypeID = tokenType
	accessToken.ExpiresAt = user.CreatedAt.Add(time.Hour * 24)
	accessToken.Hash, err = auth.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}
	return &accessToken, nil
}
