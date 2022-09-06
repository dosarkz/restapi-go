package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"rest/app/validators"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	"rest/tests"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type ResponseBodyUser struct {
	Data models.User `json:"data"`
}

// type ResponseBodyUsers struct {
//	Data []models.User `json:"data"`
//}

func TestCreateUser(t *testing.T) {
	email := faker.Email()
	password := faker.Password()
	newUser := forms.NewUser{
		Email:    email,
		Password: password,
		Phone:    faker.Phonenumber(),
		Name:     "AdminQlt",
		StatusID: models.Activated,
	}

	validate := validator.New()
	validators.RegisterCustomValidations(validate, tests.GetDb())
	err := validate.Struct(newUser)
	var errors []string

	if err != nil {
		fmt.Println(err.(validator.ValidationErrors))
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		t.Fatal(errors)
		return
	}

	// DO NOT DELETE WILL BE USED IN THE FUTURE

	// admin, err := config.GetRepositories().AdminRepo.Create(newUser)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// admin = admin

	body := forms.NewUser{
		Name:     faker.Name(),
		Phone:    faker.Phonenumber(),
		Password: faker.Password(),
		Email:    faker.Email(),
		StatusID: 1,
	}
	payloadBuf := new(bytes.Buffer)
	err = json.NewEncoder(payloadBuf).Encode(body)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", "/user", payloadBuf)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	userController := NewController(userRepo)
	handler := http.HandlerFunc(userController.Create)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	user, err := userRepo.ShowByEmail(body.Email)
	if err != nil {
		t.Fatal(err)
	}

	response := ResponseBodyUser{}
	response.Data = models.User{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		return
	}
	require.Equal(t, user.ID, response.Data.ID)
	require.Equal(t, user.Email, response.Data.Email)
	require.Equal(t, user.Name, response.Data.Name)
	require.Equal(t, user.Phone, response.Data.Phone)
}

func TestShowUser(t *testing.T) {
	// DO NOT DELETE WILL BE USED IN THE FUTURE

	// admin, err := config.GetRepositories().AdminRepo.Create(newUser)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// admin = admin

	newUser := forms.NewUser{
		Name:     faker.Name(),
		Phone:    faker.Phonenumber(),
		Password: faker.Password(),
		Email:    faker.Email(),
		StatusID: 1,
	}
	validators.SetValidator(validate)
	_, err := validators.Validate(newUser)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = userRepo.Create(newUser)
	if err != nil {
		t.Fatal(err)
	}

	user, err := userRepo.ShowByEmail(newUser.Email)
	if err != nil {
		t.Fatal(err)
	}

	idString := strconv.Itoa(int(user.ID))

	req, err := http.NewRequest("GET", "/user/"+idString, nil)
	vars := map[string]string{
		"id": idString,
	}
	req = mux.SetURLVars(req, vars)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	userController := NewController(userRepo)
	handler := http.HandlerFunc(userController.Show)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := ResponseBodyUser{}
	response.Data = models.User{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		return
	}
	require.Equal(t, user.ID, response.Data.ID)
	require.Equal(t, user.Email, response.Data.Email)
	require.Equal(t, user.Name, response.Data.Name)
	require.Equal(t, user.Phone, response.Data.Phone)
}

func TestListUser(t *testing.T) {
	// DO NOT DELETE WILL BE USED IN THE FUTURE

	// admin, err := config.GetRepositories().AdminRepo.Create(newUser)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// admin = admin

	body := forms.NewUser{
		Name:     faker.Name(),
		Phone:    faker.Phonenumber(),
		Password: faker.Password(),
		Email:    faker.Email(),
		StatusID: 1,
	}
	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(body)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", "/user", payloadBuf)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	userController := NewController(userRepo)
	handler := http.HandlerFunc(userController.Create)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// bodyList := ""
	// payloadBuf = new(bytes.Buffer)
	// json.NewEncoder(payloadBuf).Encode(bodyList)
	// reqList, err := http.NewRequest("GET", "/users", payloadBuf)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// rr = httptest.NewRecorder()
	// userController = NewController(userRepo)
	// handler = http.HandlerFunc(userController.List)
	// handler.ServeHTTP(rr, reqList)

	// users, err := userRepo.List()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// response := ResponseBodyUsers{}
	// response.Data = []models.User{}
	// err = json.Unmarshal(rr.Body.Bytes(), &response)
	// if err != nil {
	// 	return
	// }

	// require.Equal(t, len(users), len(response.Data))
}

func TestUpdateUser(t *testing.T) {
	//email := faker.Email()
	//password := faker.Password()
	//newUser := forms.NewUser{
	//	Email:    email,
	//	Password: password,
	//	Phone:    faker.Phonenumber(),
	//	Name:     "AdminQlt",
	//	StatusID: models.Activated,
	//}
	//
	//validate := validator.New()
	//validators.RegisterCustomValidations(validate, tests.GetDb())
	//err := validate.Struct(newUser)
	//var errors []string
	//
	//if err != nil {
	//	fmt.Println(err.(validator.ValidationErrors))
	//	for _, e := range err.(validator.ValidationErrors) {
	//		errors = append(errors, e.Error())
	//	}
	//
	//	t.Fatal(errors)
	//	return
	//}

	// DO NOT DELETE WILL BE USED IN THE FUTURE

	// admin, err := config.GetRepositories().AdminRepo.Create(newUser)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// admin = admin

	body := forms.NewUser{
		Name:     faker.Name(),
		Phone:    faker.Phonenumber(),
		Password: faker.Password(),
		Email:    faker.Email(),
		StatusID: 1,
	}

	validators.SetValidator(validate)
	_, err := validators.Validate(body)

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = userRepo.Create(body)
	if err != nil {
		t.Fatal(err)
	}

	updatedEmail := "Updated" + faker.Email()

	bodyUpdate := forms.UpdateUser{
		Name:     faker.Name(),
		Phone:    faker.Phonenumber(),
		Password: faker.Password(),
		Email:    updatedEmail,
		StatusID: 1,
	}
	reqBodyBytes := new(bytes.Buffer)
	err = json.NewEncoder(reqBodyBytes).Encode(bodyUpdate)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/auth/update", bytes.NewReader(reqBodyBytes.Bytes()))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	userController := NewController(userRepo)
	handler := http.HandlerFunc(userController.Create)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := ResponseBodyUser{}
	response.Data = models.User{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		return
	}

	userUpdated, err := userRepo.ShowById(response.Data.ID)
	if err != nil {
		log.Println(err)
		return
	}

	require.Equal(t, response.Data.ID, userUpdated.ID)
	require.Equal(t, bodyUpdate.Email, userUpdated.Email)
}

func TestDeleteUser(t *testing.T) {
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

	user, err := userRepo.Create(newUser)
	if err != nil {
		t.Fatal(err)
	}

	// DO NOT DELETE WILL BE USED IN THE FUTURE

	// admin, err := config.GetRepositories().AdminRepo.Create(newUser)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// admin = admin

	idString := strconv.Itoa(int(user.ID))

	req, err := http.NewRequest("DELETE", "/user/"+idString, nil)
	vars := map[string]string{
		"id": idString,
	}
	req = mux.SetURLVars(req, vars)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	userController := NewController(userRepo)
	handler := http.HandlerFunc(userController.Delete)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := ResponseBodyUser{}
	response.Data = models.User{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		return
	}
	require.Equal(t, user.ID, response.Data.ID)
	require.Equal(t, "", response.Data.Email)
	require.Equal(t, "", response.Data.Name)
	require.Equal(t, "", response.Data.Phone)
}
