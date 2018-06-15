package usrlogin

import (
	"database/sql"

	"github.com/govinda-attal/user-auth/internal/provider"
	"github.com/govinda-attal/user-auth/pkg/core/status"
)

// NewAuthenticateSrv ...
func NewAuthenticateSrv() Authenticater {
	db := provider.GetSvc(provider.SvcUserStore).(*sql.DB)
	return &authenticateSrv{db: db}
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

type authenticateSrv struct {
	db *sql.DB
}

func (as *authenticateSrv) Authenticate(rq *AutheticateRq) (*AutheticateRs, error) {
	return nil, nil
}
