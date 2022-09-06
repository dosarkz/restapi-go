package controllers

import (
	"encoding/json"
	"rest/app/helpers/auth"
	"rest/app/helpers/responses"
	"rest/app/validators"
	"rest/domain/user/forms"
	"rest/domain/user/repositories"
	"rest/router/lang"
	"net/http"
	"time"
)

type AuthController struct {
	repo     repositories.AuthRepository
	userRepo repositories.UserRepository
}

func NewAuthController(repo repositories.AuthRepository, userRepo repositories.UserRepository) *AuthController {
	return &AuthController{repo, userRepo}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginForm forms.LoginForm
	_ = json.NewDecoder(r.Body).Decode(&loginForm)

	errsVal, errVal := validators.Validate(loginForm)
	if errVal != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.validation_error"), errsVal)
		return
	}
	login, errRepo := c.repo.Login(&loginForm, c.userRepo)
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}

	responses.ResponseResource(w, login)
}

func (c *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	claims, errRequest := auth.ValidateRequest(r)
	if errRequest != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.req_val_error"), nil)
		return
	}
	claims.StandardClaims.IssuedAt = time.Now().Unix()
	claims.StandardClaims.ExpiresAt = time.Now().Add(30 * 24 * time.Hour).Unix()
	tokenFromClaims, errToken := auth.GenerateTokenFromClaims(claims)
	if errToken != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.token_error"), nil)
		return
	}
	if _, errRepo := c.repo.Logout(r); errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, tokenFromClaims)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	logout, errRepo := c.repo.Logout(r)
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, logout)
}
