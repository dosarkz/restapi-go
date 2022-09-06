package list

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	role "rest/domain/role/models"
	"log"
	"strings"
)

type Role struct {
	DB *gorm.DB
}

func (d Role) Tag() string {
	return "role"
}

func (d Role) Validate(fl validator.FieldLevel) bool {
	arr := strings.Split(fl.Param(), "-")
	if len(arr) < 2 {
		return false
	}
	var count int64
	if err := d.DB.Table(arr[0]).Where(arr[1]+" = ? AND role_id = ?", fl.Field().Uint(), role.RoleManagerInt).Count(&count).Error; err != nil {
		return false
	}
	return count == 1
}

func (d Role) Translation() *CustomTrans {
	return &CustomTrans{
		Translation: "Вы можете указать только пользователя с ролью менеджер",
		Override:    true,
		CustomRegisFunc: func(ut ut.Translator) error {
			switch ut.Locale() {
			case "en":
				if err := ut.Add("role", "You can only specify a user with the manager role", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			case "ru":
				if err := ut.Add("role", "Вы можете указать только пользователя с ролью менеджер", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			}
			return nil
		},
	}
}
