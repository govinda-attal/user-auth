package usrlogin

import (
	"database/sql"

	"github.com/govinda-attal/user-auth/pkg/core/status"
	"github.com/govinda-attal/user-auth/provider"
)

// NewAuthenticateSrv ...
func NewAuthenticateSrv() Authenticater {
	usrStore := provider.GetSvc(provider.SvcUserStore).(*sql.DB)
	return &authenticateSrv{baseSrv{usrstore: usrStore}}
}

// NewVerifySrv ...
func NewVerifySrv() Verifier {
	usrStore := provider.GetSvc(provider.SvcUserStore).(*sql.DB)
	return &verifySrv{baseSrv{usrstore: usrStore}}
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
