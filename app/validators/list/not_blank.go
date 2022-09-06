package list

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

type NotBlank struct {
}

func (d NotBlank) Tag() string {
	return "notBlank"
}

func (d NotBlank) Validate(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func (d NotBlank) Translation() *CustomTrans {
	return &CustomTrans{
		Translation: "Запрещено использовать только пробелы",
		Override:    true,
		CustomRegisFunc: func(ut ut.Translator) error {
			switch ut.Locale() {
			case "en":
				if err := ut.Add("notBlank", "It is forbidden to use only spaces", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			case "ru":
				if err := ut.Add("notBlank", "Запрещено использовать только пробелы", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			}
			return nil
		},
	}
}
