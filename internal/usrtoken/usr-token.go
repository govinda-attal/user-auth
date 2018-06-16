package usrtoken

import (
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/govinda-attal/user-auth/pkg/core/status"
)

// GetBearerToken ...
func GetBearerToken(header string) (string, error) {
	if header == "" {
		return "", status.ErrBadRequest.WithMessage("An authorization header is required")
	}
	token := strings.Split(header, " ")
	if len(token) != 2 {
		return "", status.ErrBadRequest.WithMessage("Malformed bearer token")
	}
	return token[1], nil
}

// SignJwt ...
func SignJwt(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// VerifyJwt ...
func VerifyJwt(token string, secret string) (map[string]interface{}, error) {
	jwToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.ErrInternal.WithMessage("There was an internal error")
		}
		return []byte(secret), nil
	})
	if err != nil {
		if errSvc, ok := err.(status.ErrServiceStatus); ok {
			return nil, errSvc
		}
		return nil, status.ErrInternal.WithMessage(err.Error())
	}
	if !jwToken.Valid {
		return nil, status.ErrBadRequest.WithMessage("Invalid authorization token")
	}
	return jwToken.Claims.(jwt.MapClaims), nil
}
