package validators

import (
	"github.com/go-playground/validator/v10"
	"rest/app/validators/list"
)

type CustomValidators struct {
	validators []ValidatorInterface
}

type ValidatorInterface interface {
	Validate(fl validator.FieldLevel) bool
	Tag() string
	Translation() *list.CustomTrans
}

func GetValidators() *CustomValidators {
	return &CustomValidators{
		validators: []ValidatorInterface{
			list.Exist{DB: DB},
			list.DoesntExist{DB: DB},
			list.NotBlank{},
			list.Lowercase{},
			list.StandardFloat{},
			list.Role{DB: DB},
		},
	}
}
