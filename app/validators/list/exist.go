package list

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	googleUuid "github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"reflect"
	"strings"
)

type Exist struct {
	DB *gorm.DB
}

func (d Exist) Tag() string {
	return "exists"
}

func (d Exist) existDeleteAt(table string) bool {
	result := map[string]interface{}{}
	if err := d.DB.Table("information_schema.columns").Select("column_name").Where("table_name = ? AND column_name = 'deleted_at'", table).Take(&result).Error; err != nil {
		return false
	}
	return true
}

func (d Exist) Validate(fl validator.FieldLevel) bool {
	arr := strings.Split(fl.Param(), "-")
	if len(arr) < 2 {
		return false
	}
	var count int64
	if !d.existDeleteAt(arr[0]) {
		if reflect.Slice == fl.Field().Kind() {
			v := fl.Field().Slice(0, fl.Field().Len())
			arrOfIds := []uint{}

			for _, v := range v.Interface().([]uint) {
				arrOfIds = append(arrOfIds, v)
			}
			if err := d.DB.Table(arr[0]).Where(arr[1]+" IN ?", arrOfIds).Count(&count).Error; err != nil {
				return false
			}
			if count == int64(len(arrOfIds)) {
				return true
			} else {
				return false
			}
		}

		if reflect.Uint == fl.Field().Kind() {
			if err := d.DB.Table(arr[0]).Where(arr[1]+" = ?", fl.Field().Uint()).Count(&count).Error; err != nil {
				return false
			}
		} else {
			if err := d.DB.Table(arr[0]).Where(arr[1]+" = ?", strings.TrimSpace(fl.Field().String())).Count(&count).Error; err != nil {
				return false
			}
		}
		return count > 0
	}

	if reflect.Slice == fl.Field().Kind() {
		v := fl.Field().Slice(0, fl.Field().Len())

		switch v.Interface().(type) {
		case []string:
			arrOfIds := []uuid.UUID{}
			var a = v.Interface().([]string)
			for i := 0; i < len(a); i++ {
				arrOfIds = append(arrOfIds, uuid.FromStringOrNil(a[i]))
			}
			if err := d.DB.Table(arr[0]).Where("deleted_at IS NULL AND "+arr[1]+" IN ?", arrOfIds).Count(&count).Error; err != nil {
				return false
			}
			if count == int64(len(arrOfIds)) {
				return true
			} else {
				return false
			}
		case []uint:
			arrOfIds := []uint{}
			for _, v := range v.Interface().([]uint) {
				arrOfIds = append(arrOfIds, v)
			}
			if err := d.DB.Table(arr[0]).Where("deleted_at IS NULL AND "+arr[1]+" IN ?", arrOfIds).Count(&count).Error; err != nil {
				return false
			}
			if count == int64(len(arrOfIds)) {
				return true
			} else {
				return false
			}
		case []googleUuid.UUID:
			arrOfIds := []googleUuid.UUID{}
			for _, v := range v.Interface().([]googleUuid.UUID) {
				arrOfIds = append(arrOfIds, v)
			}
			if err := d.DB.Table(arr[0]).Where("deleted_at IS NULL AND "+arr[1]+" IN ?", arrOfIds).Count(&count).Error; err != nil {
				return false
			}
			if count == int64(len(arrOfIds)) {
				return true
			} else {
				return false
			}
		}
	} else if reflect.Array == fl.Field().Kind() {
		if err := d.DB.Table(arr[0]).Where("deleted_at IS NULL AND "+arr[1]+" = ?", fl.Field().Interface().(googleUuid.UUID)).Count(&count).Error; err != nil {
			return false
		}
	}
	if reflect.Uint == fl.Field().Kind() {
		if err := d.DB.Table(arr[0]).Where("deleted_at IS NULL AND "+arr[1]+" = ?", fl.Field().Uint()).Count(&count).Error; err != nil {
			return false
		}
	} else if reflect.String == fl.Field().Kind() {
		if err := d.DB.Table(arr[0]).Where("deleted_at IS NULL AND "+arr[1]+" = ?", strings.TrimSpace(fl.Field().String())).Count(&count).Error; err != nil {
			return false
		}
	}

	return count > 0
}

func (d Exist) Translation() *CustomTrans {
	return &CustomTrans{
		Translation: "Запись с таким значением не существует",
		Override:    true,
		CustomRegisFunc: func(ut ut.Translator) error {
			switch ut.Locale() {
			case "en":
				if err := ut.Add("exists", "The record does not exists", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			case "ru":
				if err := ut.Add("exists", "Запись с таким значением не существует", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			}
			return nil
		},
	}
}
