package usrlogin

import (
	"database/sql"

	"github.com/govinda-attal/user-auth/internal/provider"
	"github.com/govinda-attal/user-auth/pkg/core/status"
)



// NewVerifySrv ...
func NewVerifySrv() Verifier {
	db := provider.GetSvc(provider.SvcUserStore).(*sql.DB)
	return &verifySrv{db: db}}
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




type verifySrv struct {
	db *sql.DB
}


func (vs *verifySrv) Verify(rq *VerificationRq) (*VerificationRs, error) {
	return nil, nil
}
