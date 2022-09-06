package list

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"unicode"
)

type Lowercase struct {
}

func (d Lowercase) Tag() string {
	return "lowercase"
}

func (d Lowercase) Validate(fl validator.FieldLevel) bool {
	field := []rune(fl.Field().String())
	var result = false
	for _, v := range field {
		if unicode.IsLower(v) == true {
			result = true
		} else {
			return false
		}
	}
	return result
}

func (d Lowercase) Translation() *CustomTrans {
	return &CustomTrans{
		Translation: "Допускается буквы только с нижним регистром",
		Override:    true,
		CustomRegisFunc: func(ut ut.Translator) error {
			switch ut.Locale() {
			case "en":
				if err := ut.Add("lowercase", "Only lowercase letters are allowed", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			case "ru":
				if err := ut.Add("lowercase", "Допускается буквы только с нижним регистром", true); err != nil {
					log.Println(err.Error())
					return nil
				}
			}
			return nil
		},
	}
}
