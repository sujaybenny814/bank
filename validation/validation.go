package validation

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/en"
)

func CheckValidation(inputModel interface{}) error {

	validate := validator.New()
	uni := ut.New(en.New())
	trans, _ := uni.GetTranslator("en")

	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}
	err = validate.Struct(inputModel)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
