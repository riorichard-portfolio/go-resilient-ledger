package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

func New() (*Validator, error) {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		jsonName := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if jsonName != "" && jsonName != "-" {
			return jsonName
		}
		queryName := strings.SplitN(fld.Tag.Get("query"), ",", 2)[0]
		if queryName != "" && queryName != "-" {
			return queryName
		}
		mapstructureName := strings.SplitN(fld.Tag.Get("mapstructure"), ",", 2)[0]
		if mapstructureName != "" && mapstructureName != "-" {
			return mapstructureName
		}
		return fld.Name
	})
	idLocale := id.New()
	uniTrans := ut.New(idLocale, idLocale)
	trans, found := uniTrans.GetTranslator("id")
	if !found {
		trans = uniTrans.GetFallback()
	}
	err := id_translations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		return nil, err
	}
	return &Validator{
		Validate:     v,
		Translator: trans,
	}, nil
}
