package usrlogin

import (
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	"github.com/govinda-attal/user-auth/internal/config"
	"github.com/govinda-attal/user-auth/internal/crypto"
	"github.com/govinda-attal/user-auth/internal/dbtx"
	"github.com/govinda-attal/user-auth/internal/provider"
	"github.com/govinda-attal/user-auth/internal/usrtoken"
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

	uid, email, ustatus, err := as.checkUserPwd(rq.Username, rq.Password)
	if err != nil {
		return nil, err
	}

	expireToken := time.Now().Add(time.Hour * 12).Unix()

	claims := make(map[string]interface{})
	claims["uid"] = uid
	claims["username"] = rq.Username
	claims["password"] = rq.Password
	claims["email"] = email
	claims["status"] = ustatus
	claims["StandardClaims"] = jwt.StandardClaims{
		ExpiresAt: expireToken,
		Issuer:    "user-auth",
	}
	tokenString, err := usrtoken.SignJwt(claims, viper.GetString(config.JwtSecret))
	if err != nil {
		return nil, status.ErrInternal.WithMessage(err.Error())
	}
	success := status.New(status.Success)
	return &AutheticateRs{Status: &success, Token: tokenString}, nil
}

func (as *authenticateSrv) checkUserPwd(username, password string) (string, string, string, error) {
	var uid, hpwd, email, ustatus string
	findUsrStmt := `SELECT 
		user_id, email, password, status 
		FROM USERS.ACCOUNT 
		WHERE username = $1`

	updUsrLoginStmt := `UPDATE 
		USERS.ACCOUNT SET last_login = now() 
		WHERE user_id = $1`

	db := as.db

	err := dbtx.WithTransaction(db, func(tx dbtx.Transaction) error {
		err := tx.QueryRow(
			findUsrStmt,
			username).Scan(&uid, &email, &hpwd, &ustatus)
		if err != nil {
			return status.ErrInternal.WithMessage(err.Error())
		}
		if ok := crypto.ComparePasswords(hpwd, password); !ok {
			return status.ErrUnauhtorized
		}

		_, err = tx.Exec(
			updUsrLoginStmt,
			uid)
		return err
	})

	if err != nil {
		return "", "", "", err
	}
	return uid, email, ustatus, nil
}
