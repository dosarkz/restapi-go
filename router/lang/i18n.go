package lang

import (
	"embed"
	"github.com/fitv/go-i18n"
	"log"
)

//go:embed locales/*.yml
var fs embed.FS

var translator *i18n.I18n

func init() {
	t, err := i18n.New(fs, "locales")
	if err != nil {
		log.Println(err)
	}
	t.SetDefaultLocale("ru")
	translator = t
}

func GetTranslator() *i18n.I18n {
	return translator
}
