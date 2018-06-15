package handler

import (
	"encoding/json"
	"net/http"

	"github.com/govinda-attal/user-auth/pkg/core/status"
	"github.com/govinda-attal/user-auth/pkg/usronboard"
)

// RegisterUser ...
func RegisterUser(w http.ResponseWriter, r *http.Request) error {
	regRq := &usronboard.RegistrationRq{}
	if err := json.NewDecoder(r.Body).Decode(regRq); err != nil {
		return status.ErrBadRequest.WithMessage(err.Error())
	}
	srv := usronboard.NewRegistererSrv()
	regRs, err := srv.Register(regRq)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&regRs)
	return nil
}

// ConfirmUser ...
func ConfirmUser(w http.ResponseWriter, r *http.Request) error {
	confRq := &usronboard.ConfirmationRq{}
	if err := json.NewDecoder(r.Body).Decode(confRq); err != nil {
		return status.ErrBadRequest.WithMessage(err.Error())
	}
	srv := usronboard.NewConfirmerSrv()
	confRs, err := srv.Confirm(confRq)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&confRs)
	return nil
}

func validateRegRq(rq *usronboard.RegistrationRq) error {
	return nil
}

func validateConfRq(rq *usronboard.ConfirmationRq) error {
	return nil
}
