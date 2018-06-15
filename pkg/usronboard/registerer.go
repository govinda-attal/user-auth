package usronboard

import (
	"database/sql"
	"time"

	"github.com/govinda-attal/user-auth/internal/config"
	"github.com/govinda-attal/user-auth/internal/crypto"
	"github.com/govinda-attal/user-auth/internal/dbtx"
	"github.com/govinda-attal/user-auth/internal/provider"
	"github.com/govinda-attal/user-auth/pkg/core/status"
	"github.com/spf13/viper"
)

// NewRegistererSrv ...
func NewRegistererSrv() Registerer {
	db := provider.GetSvc(provider.SvcUserStore).(*sql.DB)
	return &registererSrv{db: db}
}

// RegistrationRq ...
type RegistrationRq struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Email           string `json:"email,omitEmpty"`
	Mobile          string `json:"mobile,omitEmpty"`
}

// RegistrationRs ...
type RegistrationRs struct {
	Status *status.ServiceStatus `json:"status"`
	Token  string                `json:"token"`
}

// Registerer ...
type Registerer interface {
	Register(rq *RegistrationRq) (*RegistrationRs, error)
}

type registererSrv struct {
	db *sql.DB
}

func (rSrv *registererSrv) Register(rq *RegistrationRq) (*RegistrationRs, error) {
	uid, token, err := rSrv.addUserAccount(
		rq.UserName, rq.Password, rq.Email,
	)
	if err != nil {
		return nil, err
	}
	_ = uid
	success := status.New(status.Success)
	return &RegistrationRs{Status: &success, Token: token}, nil
}

func (rSrv *registererSrv) addUserAccount(username, password, email string) (string, string, error) {
	var uid, token string
	db := rSrv.db
	newAcctStmt := `INSERT INTO
		USERS.ACCOUNT (username, password, email)
		VALUES
		($1, $2, $3)
		RETURNING user_id`
	newAcctConfStmt := `INSERT INTO
		USERS.ACCT_CONFIRMATION (user_id, token, expires_on)
		VALUES
		($1, $2, $3)`

	err := dbtx.WithTransaction(db, func(tx dbtx.Transaction) error {
		err := tx.QueryRow(
			newAcctStmt,
			username,
			crypto.GenerateHash(password),
			email).Scan(&uid)
		if err != nil {
			return err
		}
		ed := viper.GetDuration(config.ConfirmExpiryDuration)
		ts := time.Now().Local()
		token = crypto.GenerateHash(uid + ts.String())
		_, err = tx.Exec(
			newAcctConfStmt,
			uid,
			token,
			ts.Add(ed))
		return err
	})

	if err != nil {
		errMsg := status.ErrInternal.WithMessage(err.Error())
		return "", "", errMsg
	}
	return uid, token, nil
}
