package handler

import (
	"encoding/json"
	"net/http"

	"github.com/govinda-attal/user-auth/pkg/core/status"
	"github.com/govinda-attal/user-auth/pkg/usrlogin"
)

// AuthenticateUser ...
func AuthenticateUser(w http.ResponseWriter, r *http.Request) error {
	authRq := &usrlogin.AutheticateRq{}
	if err := json.NewDecoder(r.Body).Decode(authRq); err != nil {
		return status.ErrBadRequest.WithMessage(err.Error())
	}
	srv := usrlogin.NewAuthenticateSrv()
	authRs, err := srv.Authenticate(authRq)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&authRs)
	return nil
}

// VerifyUser ...
func VerifyUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
