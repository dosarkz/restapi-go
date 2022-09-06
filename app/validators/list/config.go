package list

import "github.com/go-playground/validator/v10"

type CustomTrans struct {
	Translation     string
	Override        bool
	CustomRegisFunc validator.RegisterTranslationsFunc
	CustomTransFunc validator.TranslationFunc
}
