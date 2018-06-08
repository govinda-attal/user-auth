package usronboard

import (
	"database/sql"

	"github.com/govinda-attal/user-auth/pkg/core/status"
)

// NewRegistererSrv ...
func NewRegistererSrv(usrstore *sql.DB) Registerer {
	return &registereSrv{baseSrv{usrstore: usrstore}}
}

// NewConfirmerSrv ...
func NewConfirmerSrv(usrstore *sql.DB) Confirmer {
	return &confirmerSrv{baseSrv{usrstore: usrstore}}
}

// RegistrationRq ...
type RegistrationRq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitEmpty"`
	Mobile   string `json:"mobile,omitEmpty"`
}

// RegistrationRs ...
type RegistrationRs struct {
	Status *status.ServiceStatus `json:"status"`
}

// ConfirmationRq ...
type ConfirmationRq struct {
	Token string `json:"token"`
	Code  string `json:"code,omitEmpty"`
}

// ConfirmationRs ...
type ConfirmationRs struct {
	Status *status.ServiceStatus `json:"status"`
	Token  string                `json:"token,omitEmpty"`
}

// Registerer ...
type Registerer interface {
	Register(rq *RegistrationRq) (*RegistrationRs, error)
}

// Confirmer ...
type Confirmer interface {
	Confirm(rq *ConfirmationRq) (*ConfirmationRs, error)
}

type baseSrv struct {
	usrstore *sql.DB
}

type registereSrv struct {
	baseSrv
}

type confirmerSrv struct {
	baseSrv
}

func (rSrv *registereSrv) Register(rq *RegistrationRq) (*RegistrationRs, error) {
	return nil, nil
}

func (cSrv *confirmerSrv) Confirm(rq *ConfirmationRq) (*ConfirmationRs, error) {
	return nil, nil
}
