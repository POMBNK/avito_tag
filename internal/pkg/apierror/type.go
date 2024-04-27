package apierror

import (
	"encoding/json"
	"fmt"
)

type ApiError struct {
	err     error
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func (e *ApiError) Unwrap() error {
	return e.err
}

func (e *ApiError) Marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return bytes
}

func New(code int, message string, err error) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
		err:     err,
	}
}

func Newf(code int, message string, err error, args ...interface{}) *ApiError {
	return New(code, fmt.Sprintf(message, args...), err)
}

func internalError(err error) *ApiError {
	return New(500, "Internal server error", err)
}
