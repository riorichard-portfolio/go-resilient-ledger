package validator

import(
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validate *validator.Validate
	Translator ut.Translator
}