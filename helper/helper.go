package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func Trim(value string) string {
	return strings.TrimSpace(value)
}

func ParseBody[T any](r *http.Request) T {
	var model T
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		json.Unmarshal(body, &model)
	}
	return model
}

func SendResponse(w http.ResponseWriter, statusCode int, payload map[string]any) {
	w.WriteHeader(statusCode)
	payload["status_code"] = statusCode
	b, _ := json.Marshal(payload)
	w.Write(b)
}

func SendData(w http.ResponseWriter, statusCode int, data any) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"data":    data,
	})
}

func SendDataMessage(w http.ResponseWriter, statusCode int, data any, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func SendDataMessageFailed(w http.ResponseWriter, statusCode int, data any, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": false,
		"message": message,
		"data":    data,
	})
}

func SendMessage(w http.ResponseWriter, statusCode int, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"message": message,
	})
}

func SendMessageFail(w http.ResponseWriter, statusCode int, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": false,
		"message": message,
	})
}

func GetLimitOffset(r *http.Request) (limit int, offset int) {
	limit = ParseToInt(r.URL.Query().Get("limit"))
	offset = ParseToInt(r.URL.Query().Get("offset"))
	if limit == 0 {
		limit = 10
	}
	return
}

func ParseToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func FieldErrors(err error) map[string]string {
	errorMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrors {
			fieldName := ve.Field()
			tagName := ve.Tag()
			paramValue := ve.Param()

			switch tagName {
			case "required":
				errorMap[fieldName] = fmt.Sprintf("%s is required", fieldName)
			case "unique":
				errorMap[fieldName] = fmt.Sprintf("%s must be %s", fieldName, tagName)
			case "url":
				errorMap[fieldName] = fmt.Sprintf("%s must be %s example www.google.com", fieldName, tagName)
			case "min":
				errorMap[fieldName] = fmt.Sprintf("%s must be at least %s characters", fieldName, paramValue)
			case "max":
				errorMap[fieldName] = fmt.Sprintf("%s must not exceed %s characters", fieldName, paramValue)
			case "numeric":
				errorMap[fieldName] = fmt.Sprintf("%s must be numeric", fieldName)
			case "number":
				errorMap[fieldName] = fmt.Sprintf("%s must be number", fieldName)
			case "alpha":
				errorMap[fieldName] = fmt.Sprintf("%s must contain only alphabetic characters", fieldName)
			case "alphanum":
				errorMap[fieldName] = fmt.Sprintf("%s must be alphanumeric", fieldName)
			case "len":
				errorMap[fieldName] = fmt.Sprintf("%s must be exactly %s characters", fieldName, paramValue)
			case "eq":
				errorMap[fieldName] = fmt.Sprintf("%s must be %s", fieldName, paramValue)
			case "oneof":
				errorMap[fieldName] = fmt.Sprintf("%s must be one of %s", fieldName, paramValue)
			default:
				errorMap[fieldName] = fmt.Sprintf("%s is invalid", fieldName)
			}
		}
	} else {
		errorMap["error"] = err.Error()
	}

	return errorMap
}

const (
	StatusInternalServer = "Internal Server Error"
	StatusBadRequest     = "Bad Request"
	StatusSuccess        = "Success"
	StatusUnauthorized   = "Unauthorized"
)

var (
	ErrNotFound                 = errors.New("data not found")
	ErrUsernameAlreadyExist     = errors.New("username already exist")
	ErrEmailAlreadyExist        = errors.New("email already exist")
	ErrPhoneNumberAlreadyExist  = errors.New("phone number already exist")
	ErrCardNumberAlreadyExist   = errors.New("card number already exist")
	ErrCvvAlreadyExist          = errors.New("cvv already exist")
	ErrMinimalTransaction       = errors.New("amount must be greater than 10000")
	ErrNotEnoughBalance         = errors.New("not enough balance")
	ErrFailedCreate             = errors.New("failed to create data")
	ErrFailedCreateToken        = errors.New("failed to create token")
	ErrInvalidParseToken        = errors.New("invalid to parse token")
	ErrInvalidTokenMapclaims    = errors.New("invalid token mapclaims ")
	ErrInvalidTokenStringMethod = errors.New("invalid token string method")
	ErrInvalidExtension         = errors.New("extention is not allowed")
	ErrInvalidUsernamePassword  = errors.New("invalid username password")
	ErrTokenNotProvided         = errors.New("token not provided")
	ErrFailedUpdate             = errors.New("failed to update data")
	ErrFailedUpload             = errors.New("failed to upload data")
	ErrFailedDelete             = errors.New("failed to delete data")
	ErrTitleAlreadyExist        = errors.New("title already exist")
	ErrInvalidPage              = errors.New("invalid page")
	ErrInvalidPerPage           = errors.New("invalid per page")
)
