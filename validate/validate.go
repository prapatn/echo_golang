package validate

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

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
			upperCamelToSnake(err.Field()): CustomErrorMessage(err),
		}

		mergo.Merge(&errors, error) //mergo.Merge(&dest,src)
	}
	return &errors
}

func MapErrorBind(err error) *map[string]interface{} {
	echoErrs, _ := err.(*echo.HTTPError)
	customErr, _ := echoErrs.Internal.(interface{}).(*json.UnmarshalTypeError)
	errorMessage := echoErrs.Internal.Error()

	error := map[string]interface{}{
		customErr.Field: errorMessage,
	}

	return &error
}

func upperCamelToSnake(input string) string {
	var result strings.Builder

	for i, char := range input {
		if i > 0 && unicode.IsUpper(char) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(char))
	}

	return result.String()
}

func CustomErrorMessage(e validator.FieldError) string {
	field := upperCamelToSnake(e.Field())
	tag := e.Tag()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s is min %s", field, e.Param())
	case "max":
		return fmt.Sprintf("%s is max %s", field, e.Param())
	case "customUsernameValidator":
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
