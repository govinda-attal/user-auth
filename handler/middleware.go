package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"

	"github.com/govinda-attal/user-auth/internal/config"
	"github.com/govinda-attal/user-auth/internal/usrtoken"
	"github.com/govinda-attal/user-auth/pkg/core/status"
)

type key string

const (
	dt key = "decodedToken"
	// ...
)

// ValidateUserLogon ...
func ValidateUserLogon(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bearerToken, err := usrtoken.GetBearerToken(req.Header.Get("Authorization"))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		decodedToken, err := usrtoken.VerifyJwt(bearerToken, viper.GetString(config.JwtSecret))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		if decodedToken["authorized"] == true {
			ctx := context.WithValue(req.Context(), dt, decodedToken)
			req = req.WithContext(ctx)
			next(w, req)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			err := status.ErrUnauhtorized.WithMessage("2FA is required")
			json.NewEncoder(w).Encode(err)
		}
	})

}
