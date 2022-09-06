package seeders

import (
	"fmt"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	"rest/domain/user/repositories"
	"os"
)

func Run(repository repositories.UserRepository) {
	generateDefaultAdminUser(repository)
}

func generateDefaultAdminUser(repository repositories.UserRepository) {
	email := os.Getenv("ADMIN_DEMO_USER_EMAIL")
	user := forms.NewUser{
		Email:    email,
		Password: os.Getenv("ADMIN_DEMO_USER_PASSWORD"),
		Phone:    "123456789456",
		Name:     "AdminQlt",
		StatusID: models.Activated,
	}

	field, _ := repository.FindBy(email, "email")
	if field.Email != "" {
		//	fmt.Println("Admin has been already created, email: " + field.Email)
		return
	}

	u, err := repository.Create(user)
	if err != nil {
		return
	}

	fmt.Println("Default admin user created. email: " + u.Email)
}
