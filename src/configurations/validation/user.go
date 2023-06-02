package validation

import (
	"encoding/json"
	"errors"
	"github.com/caio-rds/golang-api/src/configurations/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	//Validate = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en_ := en.New()
		un := ut.New(en_, en_)
		transl, _ := un.GetTranslator("en_")
		err := enTranslations.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValErr validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid Field Type")
	} else if errors.As(validationErr, &jsonValErr) {
		var errorsCauses []rest_err.Causes

		for _, err := range validationErr.(validator.ValidationErrors) {
			errorsCauses = append(errorsCauses, rest_err.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			})
		}
		return rest_err.NewBadRequestValidationError("Invalid Request Body", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Failed to convert fields")
	}
}
