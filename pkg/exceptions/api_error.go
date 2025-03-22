package exceptions

import "net/http"

const (
	CODE_USER_NOT_FOUND              = 1
	CODE_USER_BAD_ID                 = 2
	CODE_USER_EXTERNAL_SERVICE_ERROR = 3
	CODE_USER_ALREADY_EXISTS         = 4
)

const (
	CODE_GENERIC_DB_ERROR        = 401
	CODE_REQUEST_INVALID_PAYLOAD = 400
	CODE_INTERNAL_ERROR          = 500
)

type APIError struct {
	HttpCode  int      `json:"http_code"`
	ErrorCode ErrorMsg `json:"error_code"`
}

type ErrorMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err APIError) Error() string {
	return err.ErrorCode.Message
}

var (
	ErrUserNotFound          = APIError{HttpCode: http.StatusNotFound, ErrorCode: ErrorMsg{Code: CODE_USER_NOT_FOUND, Message: "User not found"}}
	ErrUserBadId             = APIError{HttpCode: http.StatusBadRequest, ErrorCode: ErrorMsg{Code: CODE_USER_BAD_ID, Message: "Bad user id"}}
	ErrUserExternalService   = APIError{HttpCode: http.StatusServiceUnavailable, ErrorCode: ErrorMsg{Code: CODE_USER_EXTERNAL_SERVICE_ERROR, Message: "External service error"}}
	ErrUserAlreadyExists     = APIError{HttpCode: http.StatusConflict, ErrorCode: ErrorMsg{Code: CODE_USER_ALREADY_EXISTS, Message: "User already exists"}}
	ErrGenericDb             = APIError{HttpCode: http.StatusInternalServerError, ErrorCode: ErrorMsg{Code: CODE_GENERIC_DB_ERROR, Message: "Generic database error"}}
	ErrInvalidRequestPayload = APIError{HttpCode: http.StatusBadRequest, ErrorCode: ErrorMsg{Code: CODE_REQUEST_INVALID_PAYLOAD, Message: "Invalid request payload"}}
	ErrInternal              = APIError{HttpCode: http.StatusInternalServerError, ErrorCode: ErrorMsg{Code: CODE_INTERNAL_ERROR, Message: "Internal Server Error"}}
)
