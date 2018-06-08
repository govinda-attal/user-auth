package handler

import (
	"net/http"

	"github.com/govinda-attal/user-auth/provider"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	usrStore := provider.GetSvc(provider.SvcUserStore)
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {

}
