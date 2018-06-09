package usrtoken

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/govinda-attal/user-auth/internal/vars"
	"github.com/govinda-attal/user-auth/pkg/core/status"
)

type key string

const (
	dt key = "decodedToken"
	// ...
)

// ValidateUserLogon ...
func ValidateUserLogon(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	bearerToken, err := GetBearerToken(req.Header.Get("Authorization"))
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	decodedToken, err := VerifyJwt(bearerToken, vars.GetVar("jwtSecret"))
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
}
