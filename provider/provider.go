package provider

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/govinda-attal/user-auth/provider/usrstore"
)

const SvcUserStore = "svc.usrstore"

func Setup() {
	db, err := usrstore.InitStore()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault(SvcUserStore, db)
}

func GetSvc(svcName string) interface{} {
	return viper.Get(SvcUserStore)
}

func Cleanup() {
	db := viper.Get(SvcUserStore)
	if db != nil {
		err := db.(*sql.DB).Close()
		if err != nil {
			fmt.Println("Close on db connection failed!")
		}
	}
}
