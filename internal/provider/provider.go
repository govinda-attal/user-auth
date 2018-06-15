package provider

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/govinda-attal/user-auth/internal/provider/usrstore"
)

const (
	// SvcUserStore ...
	SvcUserStore = "svc.usrstore"
)

// Setup ...
func Setup() {
	db, err := usrstore.InitStore()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault(SvcUserStore, db)
}

// GetSvc ...
func GetSvc(svcName string) interface{} {
	return viper.Get(SvcUserStore)
}

// Cleanup ...
func Cleanup() {
	db := viper.Get(SvcUserStore)
	if db != nil {
		err := db.(*sql.DB).Close()
		if err != nil {
			fmt.Println("Close on db connection failed!")
		}
	}
}
