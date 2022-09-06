package validators

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	ruTrans "github.com/go-playground/validator/v10/translations/ru"
	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
	"rest/database/redis"
	"log"
)

var validate *validator.Validate
var vl *Validator

type Validator struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
}

func Init(db *gorm.DB) {
	newVl := NewValidator()
	RegisterCustomValidations(newVl.validate, db)

	for _, locale := range []string{"ru", "en"} {
		t, found := newVl.uni.GetTranslator(locale)
		if !found {
			log.Println("locale not found", locale)
		}
		err := RegisterCustomTranslations(newVl.validate, t)
		if err != nil {
			log.Println(err)
		}
	}
}

func NewValidator() *Validator {
	vl = &Validator{
		validate: validator.New(),
		uni:      ut.New(ru.New(), en.New()),
	}

	locale, _ := GetRedisValue("locale")
	if locale == "" {
		locale = "ru"
	}

	vl.SetTranslator(locale)
	return vl
}

func (v *Validator) SetTranslator(locale string) {
	trans, _ = v.uni.GetTranslator(locale)

	var err error
	switch locale {
	case "ru":
		err = ruTrans.RegisterDefaultTranslations(v.validate, trans)
	case "en":
		err = enTrans.RegisterDefaultTranslations(v.validate, trans)
	}

	if err != nil {
		log.Println(err)
	}
}

func Validate(input any) (map[string]interface{}, error) {
	errors := make(map[string]interface{})

	if GetValidator() == nil {
		validate = NewValidator().validate
	} else {
		locale, _ := GetRedisValue("locale")
		GetValidator().SetTranslator(locale)
		validate = GetValidator().validate
	}

	err := validate.Struct(input)

	if err != nil {
		var arrStr []string
		for _, e := range err.(validator.ValidationErrors) {
			errors[strcase.ToLowerCamel(e.Field())] = append(arrStr, e.Translate(trans))
		}
		return errors, err
	}
	return errors, nil
}

func SetValidator(newValid *validator.Validate) {
	validate = newValid
}

func GetValidator() *Validator {
	return vl
}

func GetRedisValue(value string) (string, error) {
	rConf := new(redis.Config)
	redisConn := redis.ConnectToRedis(rConf)
	getVal, err := redisConn.GetValue(value)
	return getVal, err
}
