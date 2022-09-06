package list

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"rest/app/helpers/type_helpers/regexes"
	"log"
	"reflect"
)

type StandardFloat struct {
}

func (d StandardFloat) Tag() string {
	return "standardFloat"
}

func (d StandardFloat) Validate(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind() == reflect.Float64 {
		return regexes.IsStandardFloat(field.Float())
	}
	return false
}

func (d StandardFloat) Translation() *CustomTrans {
	return &CustomTrans{
		Translation: "Только числа с плавающей точкой",
		Override:    true,
		CustomRegisFunc: func(ut ut.Translator) error {
			switch ut.Locale() {
			case "en":
				if err := ut.Add("standardFloat", "Only floating point numbers", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			case "ru":
				if err := ut.Add("standardFloat", "Только числа с плавающей точкой", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			}
			return nil
		},
	}
}
