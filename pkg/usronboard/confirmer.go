package usronboard

import (
	"database/sql"
	"time"

	"github.com/govinda-attal/user-auth/internal/dbtx"
	"github.com/govinda-attal/user-auth/internal/provider"
	"github.com/govinda-attal/user-auth/pkg/core/status"
)

// NewConfirmerSrv ...
func NewConfirmerSrv() Confirmer {
	db := provider.GetSvc(provider.SvcUserStore).(*sql.DB)
	return &confirmerSrv{db: db}
}

// ConfirmationRq ...
type ConfirmationRq struct {
	Token string `json:"token"`
	Code  string `json:"code,omitEmpty"`
}

// ConfirmationRs ...
type ConfirmationRs struct {
	Status *status.ServiceStatus `json:"status"`
	UserID string                `json:"userID"`
}

// Confirmer ...
type Confirmer interface {
	Confirm(rq *ConfirmationRq) (*ConfirmationRs, error)
}

type confirmerSrv struct {
	db *sql.DB
}

func (cSrv *confirmerSrv) Confirm(rq *ConfirmationRq) (*ConfirmationRs, error) {
	uid, err := cSrv.confirmAcct(rq.Token)
	if err != nil {
		return nil, err
	}
	success := status.New(status.Success)
	return &ConfirmationRs{Status: &success, UserID: uid}, nil
}

func (cSrv *confirmerSrv) confirmAcct(token string) (string, error) {
	var uid string
	confirmAcctStmt := `UPDATE 
		USERS.ACCT_CONFIRMATION SET confirmed = $2, confirmed_on = $3 
		WHERE token = $1
		RETURNING user_id`

	activateAcctStmt := `UPDATE 
		USERS.ACCOUNT SET status = 'ACTIVE' 
		WHERE user_id = $1`
	db := cSrv.db

	err := dbtx.WithTransaction(db, func(tx dbtx.Transaction) error {
		err := tx.QueryRow(
			confirmAcctStmt,
			token,
			true,
			time.Now().Local()).Scan(&uid)
		if err != nil {
			return err
		}

		_, err = tx.Exec(
			activateAcctStmt,
			uid)
		return err
	})

	if err != nil {
		errMsg := status.ErrInternal.WithMessage(err.Error())
		return "", errMsg
	}
	return uid, nil
}
