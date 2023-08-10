package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
)

func Init() *validator.Validate {
	validate := validator.New()
	// Register the custom validator function
	validate.RegisterValidation("customUsernameValidator", CustomUsernameValidator)

	return validate
}

func MapErrorValidate(err error) *map[string]interface{} {
	var errors map[string]interface{}
	for _, err := range err.(validator.ValidationErrors) {
		error := map[string]interface{}{
			err.Field(): CustomErrorMessage(err),
		}

		mergo.Merge(&errors, error) //mergo.Merge(&dest,src)
	}
	return &errors
}

func CustomErrorMessage(e validator.FieldError) string {
	field := e.Field()
	tag := e.Tag()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s is min %s", field, e.Param())
	case "max":
		return fmt.Sprintf("%s is max %s", field, e.Param())
	case "customusernamevalidator":
		return fmt.Sprintf("Custom error message for %s", field)
	case "eqfield":
		return fmt.Sprintf("%s not eqfield %s", field, e.Param())
	default:
		return fmt.Sprintf("Validation error on %s with tag %s", field, tag)
	}
}

func CustomUsernameValidator(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return len(username) >= 6
}
