package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"rest/app/helpers/env"
	"rest/app/validators"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	"rest/domain/user/repositories"
	"rest/domain/user/responses"
	"rest/tests"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var (
	validate *validator.Validate
	_        ut.Translator
	userRepo repositories.UserRepository
	authRepo repositories.AuthRepository
)

func TestMain(m *testing.M) {
	validate = validator.New()
	dbConf := new(tests.Database)
	env.HasEnvironment()
	tests.InitApp()
	tests.LoadDB(dbConf)
	db := tests.GetDb()
	newEnTranslator := en.New()
	uni := ut.New(newEnTranslator, newEnTranslator)
	_, _ = uni.GetTranslator("en")

	validators.Init(db)
	authRepo = repositories.NewAuthRepo(db)
	userRepo = repositories.NewUserRepo(db)
	code := m.Run()

	if err := tests.CloseDb(db); err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func TestLogin(t *testing.T) {
	_ = login(t)
}

func login(t *testing.T) (token string) {
	email := faker.Email()
	password := faker.Password()
	newUser := forms.NewUser{
		Email:    email,
		Password: password,
		Phone:    faker.Phonenumber(),
		Name:     "AdminQlt",
		StatusID: models.Activated,
	}
	validators.SetValidator(validate)
	_, err := validators.Validate(newUser)

	if err != nil {
		log.Fatal(err)
		return
	}

	u, err := userRepo.Create(newUser)
	if err != nil {
		t.Fatal(err)
	}

	l := forms.LoginForm{Email: email,
		Password: password}

	reqBodyBytes := new(bytes.Buffer)
	err = json.NewEncoder(reqBodyBytes).Encode(l)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "api/v1/auth/login", bytes.NewReader(reqBodyBytes.Bytes()))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	userController := NewAuthController(authRepo, userRepo)
	handler := http.HandlerFunc(userController.Login)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	response := ResponseBodyLogin{}
	response.Data = responses.UserResponse{}

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("parse response error: %v", err)
	}

	if response.Data.Token == "" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), u)
	}

	if response.Data.Email != u.Email {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), u)
	}
	token = response.Data.Token
	return
}

func TestRefreshToken(t *testing.T) {
	oldToken := login(t)
	req, err := http.NewRequest("GET", "api/v1/auth/refresh-token", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", oldToken)

	rr := httptest.NewRecorder()

	userController := NewAuthController(authRepo, userRepo)
	handler := http.HandlerFunc(userController.RefreshToken)
	//sleep for different ExpiredAt field for oldToken and newToken
	time.Sleep(1 * time.Second)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	response := ResponseBodyRefreshToken{}

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("parse response error: %v", err)
	}
	newToken := response.Data
	if oldToken == newToken {
		t.Error("refreshed token and old token can not be equal")
	}
}

type ResponseBodyLogin struct {
	Data responses.UserResponse `json:"data"`
}
type ResponseBodyRefreshToken struct {
	Data string `json:"data"`
}
