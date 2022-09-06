package models

import (
	"log"
	"rest/app/helpers/hash"
	"rest/app/helpers/type_helpers/slices"
	"rest/app/helpers/type_helpers/str"
	role "rest/domain/role/models"
	"strconv"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

const (
	Deactivated = 0
	Activated   = 1
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Phone     string         `json:"phone" gorm:"unique"`
	Password  string         `json:"password"`
	RoleID    int            `json:"role_id"`
	Role      role.Role      `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Email     string         `json:"email"`
	StatusID  int            `json:"status_id"`
}

func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	if u.Password == "" {
		return nil
	}
	u.Password, err = hash.PasswordHash(u.Password)
	return err
}

func Filter(db *gorm.DB, filters map[string]string) *gorm.DB {
	id, idOk := filters["id"]
	if idOk && str.NotEmpty(id) {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}
		db = db.Where("id = ?", uint(idInt))
	}

	name, nameOk := filters["name"]
	if nameOk && str.NotEmpty(name) {
		db = db.Where("name ILIKE ?", "%"+name+"%")
	}

	email, emailOk := filters["email"]
	if emailOk && str.NotEmpty(email) {
		db = db.Where("email ILIKE ?", "%"+email+"%")
	}

	roleID, roleIDOk := filters["roleId"]
	if roleIDOk && str.NotEmpty(roleID) {
		idInt, err := strconv.Atoi(roleID)
		if err != nil {
			log.Println(err)
		}
		db = db.Where("role_id = ?", uint(idInt))
	}

	phone, phoneOk := filters["phone"]
	if phoneOk && str.NotEmpty(phone) {
		db = db.Where("phone ILIKE ?", "%"+phone+"%")
	}

	return db
}

func Sort(db *gorm.DB, params map[string]string) *gorm.DB {

	sort, sortOk := params["sort"]
	stringSortableFields := []string{"name", "email", "phone"}
	intSortableFields := []string{"id", "roleId"}
	sortBy, sortByOk := params["sortBy"]
	sortableParams := []string{"asc", "desc"}

	if sortOk && slices.StringContain(stringSortableFields, sort) && sortByOk && slices.StringContain(sortableParams, sortBy) {
		db = db.Order("lower(" + strcase.ToSnake(sort) + ")  " + strings.ToUpper(sortBy))
		return db
	} else if sortOk && slices.StringContain(intSortableFields, sort) && sortByOk && slices.StringContain(sortableParams, sortBy) {
		db = db.Order(strcase.ToSnake(sort) + " " + strings.ToUpper(sortBy))
		return db
	}
	return db.Order("users.id desc")
}
