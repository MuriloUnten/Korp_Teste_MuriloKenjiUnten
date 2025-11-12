package internal

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"-"`
	Msg        any `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(status int, msg any) APIError {
	return APIError{
		StatusCode: status,
		Msg: msg,
	}
}

func InternalError() APIError {
	return APIError{
		StatusCode: http.StatusInternalServerError,
		Msg: "internal error",
	}
}

func BadRequest() APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Msg: "bad request",
	}
}

func InvalidRequestBody() APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Msg: "invalid request body",
	}
}

func InvalidPathIdentifier() APIError {
	return APIError{
		StatusCode: http.StatusBadRequest,
		Msg: "invalid path identifier",
	}
}

func InvalidJSONRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg: errors,
	}
}

func NotImplemented() APIError {
	return APIError{
		StatusCode: http.StatusNotImplemented,
		Msg: "Endpoint not implemented",
	}
}

func NotFound() APIError {
	return APIError{
		StatusCode: http.StatusNotFound,
		Msg: "resource not found",
	}
}
