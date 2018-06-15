package handler

import (
	"encoding/json"
	"net/http"

	"github.com/govinda-attal/user-auth/pkg/core/status"
)

// ErrorHandler ...
func ErrorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			if errSvc, ok := err.(status.ErrServiceStatus); ok {
				w.WriteHeader(errSvc.Code)
				json.NewEncoder(w).Encode(&errSvc)
				return
			}
			errSvc := status.ErrInternal.WithMessage(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&errSvc)
		}
	}
}
