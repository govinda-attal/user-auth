package usrlogin

import (
	"database/sql"

	"github.com/govinda-attal/user-auth/pkg/core/status"
)

// NewAuthenticateSrv ...
func NewAuthenticateSrv(usrstore *sql.DB) Authenticater {
	return &authenticateSrv{baseSrv{usrstore: usrstore}}
}

// NewVerifySrv ...
func NewVerifySrv(usrstore *sql.DB) Verifier {
	return &verifySrv{baseSrv{usrstore: usrstore}}
}

// AutheticateRq ...
type AutheticateRq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AutheticateRs ...
type AutheticateRs struct {
	Status *status.ServiceStatus `json:"status"`
	Token  string                `json:"token,omitEmpty"`
}

// Authenticater ...
type Authenticater interface {
	Authenticate(rq *AutheticateRq) (*AutheticateRs, error)
}

// VerificationRq ...
type VerificationRq struct {
	Token string `json:"token"`
	Code  string `json:"code"`
}

// VerificationRs ...
type VerificationRs struct {
	Status *status.ServiceStatus `json:"status"`
	Token  string                `json:"token,omitEmpty"`
}

// Verifier ...
type Verifier interface {
	Verify(rq *VerificationRq) (*VerificationRs, error)
}

type baseSrv struct {
	usrstore *sql.DB
}

type authenticateSrv struct {
	baseSrv
}

type verifySrv struct {
	baseSrv
}

func (as *authenticateSrv) Authenticate(rq *AutheticateRq) (*AutheticateRs, error) {
	return nil, nil
}

func (vs *verifySrv) Verify(rq *VerificationRq) (*VerificationRs, error) {
	return nil, nil
}
