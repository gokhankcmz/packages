package Validator

import (
	"Packages/src/api/Type/ErrorTypes"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

type ValidationErrors []error

func (errs ValidationErrors) CombineAsString() string {
	res := make([]string, len(errs))
	for i, e := range errs {
		res[i] = e.Error()
	}
	return strings.Join(res, ". ") + "."
}

func ValidateModelOrPanic(model interface{}) {
	validate := validator.New()
	err := validate.Struct(model)
	var ErrorText string
	if err != nil {
		english := en.New()
		uni := ut.New(english, english)
		trans, _ := uni.GetTranslator("en")
		_ = enTranslations.RegisterDefaultTranslations(validate, trans)
		errs := TranslateError(err, trans)
		ErrorText = ValidationErrors(errs).CombineAsString()
		panic(ErrorTypes.InvalidModel.SetPublicDetail(ErrorText))
	}
}
