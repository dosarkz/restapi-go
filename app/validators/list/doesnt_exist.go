package list

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
)

type DoesntExist struct {
	DB *gorm.DB
}

func (d DoesntExist) Tag() string {
	return "doesnt_exist"
}

func (d DoesntExist) Validate(fl validator.FieldLevel) bool {
	arr := strings.Split(fl.Param(), "-")
	if len(arr) < 2 {
		return false
	}
	var count int64
	var id uint64

	if fl.Parent().FieldByName("Id").Kind() == reflect.Uint {
		id = fl.Parent().FieldByName("Id").Uint()
	}
	if reflect.Uint == fl.Field().Kind() {
		if err := d.DB.Table(arr[0]).
			Where(arr[1]+" = ?", fl.Field().Uint()).
			Count(&count).Error; err != nil {
			return false
		}
	} else {
		if id != 0 {
			_ = d.DB.Table(arr[0]).
				Where(arr[1]+" = ? and id = ?", strings.TrimSpace(fl.Field().String()), id).Count(&count)
			if count == 1 {
				return true
			}
		}
		if err := d.DB.Table(arr[0]).Where(arr[1]+" = ?", strings.TrimSpace(fl.Field().String())).Count(&count).Error; err != nil {
			return false
		}
	}
	return count == 0
}

func (d DoesntExist) Translation() *CustomTrans {
	return &CustomTrans{
		Translation: "Запись с этим значением уже существует",
		Override:    true,
		CustomRegisFunc: func(ut ut.Translator) error {
			switch ut.Locale() {
			case "en":
				if err := ut.Add("doesnt_exist", "The record has been already exists", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			case "ru":
				if err := ut.Add("doesnt_exist", "Запись с этим значением уже существует", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			}
			return nil
		},
	}
}
