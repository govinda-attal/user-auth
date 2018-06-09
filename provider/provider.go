package provider

import (
	"log"

	"github.com/govinda-attal/user-auth/provider/usrstore"
)

const SvcUserStore = "usrstore"

var svcMap = make(map[string]interface{})


func Setup(connStr string) {
	db, err := usrstore.InitStore(connStr)
	if err != nil {
		log.Fatal(err)
	}
	svcMap["usrstore"] = db
}

func GetSvc(svcName string) interface{} {
	return svcMap[svcName]
}

