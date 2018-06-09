package status

import (
	"fmt"
	"net/http"
)

var ErrInternal = ErrServiceStatus{
	ServiceStatus{Code: http.StatusInternalServerError, Message: "Internal Server Error"},
}
var ErrNotFound = ErrServiceStatus{
	ServiceStatus{Code: http.StatusNotFound, Message: "Not Found"},
}
var ErrBadRequest = ErrServiceStatus{
	ServiceStatus{Code: http.StatusBadRequest, Message: "Bad Request"},
}
var ErrUnauhtorized = ErrServiceStatus{
	ServiceStatus{Code: http.StatusUnauthorized, Message: "Unauthorized"},
}

// ServiceStatus ...
type ServiceStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrServiceStatus ...
type ErrServiceStatus struct {
	ServiceStatus
}

func (e ErrServiceStatus) Error() string {
	return fmt.Sprintf(string(e.Code), ": ", e.Message)
}

// WithMessage ...
func (e *ErrServiceStatus) WithMessage(msg string) *ErrServiceStatus {
	e.Message = msg
	return e
}
