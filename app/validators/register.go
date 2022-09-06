package validators

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
)

var (
	DB    *gorm.DB
	trans ut.Translator
)

func RegisterCustomValidations(validate *validator.Validate, curConn *gorm.DB) {
	DB = curConn
	for _, v := range GetValidators().validators {
		if err := validate.RegisterValidation(v.Tag(), v.Validate); err != nil {
			log.Println(err)
			continue
		}
	}
}

func RegisterCustomTranslations(v *validator.Validate, trans ut.Translator) (err error) {
	for _, t := range GetValidators().validators {
		switch {
		case t.Translation().CustomTransFunc != nil && t.Translation().CustomRegisFunc != nil:
			err = v.RegisterTranslation(t.Tag(), trans, t.Translation().CustomRegisFunc, t.Translation().CustomTransFunc)
		case t.Translation().CustomTransFunc != nil && t.Translation().CustomRegisFunc == nil:
			err = v.RegisterTranslation(t.Tag(), trans, registrationFunc(t.Tag(), t.Translation().Translation,
				t.Translation().Override), t.Translation().CustomTransFunc)
		case t.Translation().CustomTransFunc == nil && t.Translation().CustomRegisFunc != nil:
			err = v.RegisterTranslation(t.Tag(), trans, t.Translation().CustomRegisFunc, translateFunc)
		default:
			err = v.RegisterTranslation(t.Tag(), trans, registrationFunc(t.Tag(), t.Translation().Translation,
				t.Translation().Override), translateFunc)
		}

		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}
		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
