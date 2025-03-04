package utils

import (
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
	}
)

var validate = validator.New()
var Logger zerolog.Logger
var ILogger zerolog.Logger
var appValidator *XValidator

func SetLogger() {
	runLogFile, _ := os.OpenFile(
        "myapp.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0664,
    )
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	logger := zerolog.New(multi).Level(zerolog.ErrorLevel).With().Timestamp().Logger()
	Logger = logger

	ilogger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	ILogger = ilogger
}

// func GetLogger() zerolog.Logger {
// 	return Logger
// }

func SetValidator() {
	v := &XValidator{
		validator: validate,
	}
	appValidator = v
}

func GetValidator() *XValidator {
	return appValidator
}

func ValidatePayload(data interface{}, c *fiber.Ctx) error {
	errs := GetValidator().Validate(data)
	if len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)
		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": errMsgs, 
			"error": fiber.ErrBadRequest.Message, 
			"statusCode": fiber.StatusBadRequest,
		})

		// return &fiber.Error{
		// 	Code: fiber.ErrBadRequest.Code,
		// 	Message: strings.Join(errMsgs, " and "),
		// }
	}

	return nil
}

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}