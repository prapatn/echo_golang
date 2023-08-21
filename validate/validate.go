package validate

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
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
	validate.RegisterValidation("customZipcodeValidator", CustomZipcodeValidator)
	validate.RegisterValidation("isValidThai", isValidThai)
	validate.RegisterValidation("empty", ValidateNotEmptySlice)

	return validate
}

func MapErrorValidate(err error) *map[string]interface{} {
	// var errors map[string]interface{}
	// for _, err := range err.(validator.ValidationErrors) {
	// 	error := map[string]interface{}{
	// 		upperCamelToSnake(err.Field()): CustomErrorMessage(err),
	// 	}

	// 	mergo.Merge(&errors, error) //mergo.Merge(&dest,src)
	// }
	// return &errors

	errors := make(map[string]interface{})
	for _, e := range err.(validator.ValidationErrors) {
		fieldNamespace := e.StructNamespace()
		// fieldName := upperCamelToSnake(e.Field())
		// errorMessage := fmt.Sprintf("Validation error on %s with tag %s", fieldName, e.Tag())
		nestedFields := strings.Split(fieldNamespace, ".")[1:]
		nestedMap := errors
		for i, nestedField := range nestedFields {
			nestedField = upperCamelToSnake(nestedField)
			if i == len(nestedFields)-1 {
				nestedMap[nestedField] = fmt.Sprintf("Validation error on %s with tag %s", nestedField, e.Tag())
			} else {
				var fields []string
				if strings.Contains(nestedField, "[") {
					fields = strings.Split(nestedField, "[")
				}

				if fields != nil {
					if _, exists := nestedMap[fields[0]]; !exists {
						nestedMap[fields[0]] = make([]map[string]interface{}, 0)
					}

					nestedErrors := nestedMap[fields[0]].([]map[string]interface{})

					if _, exists := nestedMap[nestedField]; !exists {
						nestedMap[nestedField] = make(map[string]interface{})
					}

					if nestedMap[nestedField] != nil {
						nestedError := nestedMap[nestedField].(map[string]interface{})
						nestedErrors = append(nestedErrors, nestedError)
					}

					nestedMap[fields[0]] = nestedErrors
				} else {
					if _, exists := nestedMap[nestedField]; !exists {
						nestedMap[nestedField] = make(map[string]interface{})
					}
					nestedMap = nestedMap[nestedField].(map[string]interface{})
				}
			}
		}
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

// func CustomErrorMessage(e validator.FieldError) string {
// 	field := upperCamelToSnake(e.Field())
// 	tag := e.Tag()
// 	param := upperCamelToSnake(e.Param())

// 	switch tag {
// 	case "required":
// 		return fmt.Sprintf("%s is required", field)
// 	case "min":
// 		return fmt.Sprintf("%s is min %s", field, param)
// 	case "max":
// 		return fmt.Sprintf("%s is max %s", field, param)
// 	case "customUsernameValidator":
// 		return fmt.Sprintf("Custom error message for %s", field)
// 	case "eqfield":
// 		return fmt.Sprintf("%s not eqfield %s", field, param)
// 	default:
// 		return fmt.Sprintf("Validation error on %s with tag %s", field, tag)
// 	}
// }

func CustomZipcodeValidator(fl validator.FieldLevel) bool {
	zipcode := fl.Field().String()
	return len(zipcode) == 5
}

// check thai
func isValidThai(fl validator.FieldLevel) bool {

	// text := "สวัสดียินดีต้อนรับ"
	// re := regexp.MustCompile(`^[ก-๙เ-๎]+$`)

	// if re.MatchString(text) {
	// 	fmt.Println("พบข้อความที่ประกอบด้วยตัวอักษรไทยเท่านั้น:", text)
	// } else {
	// 	fmt.Println("ไม่พบข้อความที่ประกอบด้วยตัวอักษรไทยเท่านั้น")
	// }

	value := fl.Field().String()
	var thaiNameRegex = regexp.MustCompile(`^[\p{Thai} ]+$`)

	var thaiNumberRegex = regexp.MustCompile(`[๑-๙]+`)
	matches := thaiNumberRegex.FindAllString(value, -1)

	return thaiNameRegex.MatchString(value) && (matches == nil)
}

func ValidateNotEmptySlice(fl validator.FieldLevel) bool {
	childrenValue := fl.Field()

	// Convert the field's value to an interface{}
	childrenInterface := reflect.ValueOf(childrenValue.Interface()).Interface()

	// Check the length of the interface (slice)
	childrenLen := reflect.ValueOf(childrenInterface).Len()
	return childrenLen > 0
}
