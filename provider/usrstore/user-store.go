package usrstore

import (
	"database/sql"
	"fmt"
	// ...
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitStore() (*sql.DB, error) {
	usConf := viper.GetStringMap("services.userstore")
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		usConf["host"].(string),
		usConf["port"].(int),
		usConf["username"].(string),
		usConf["password"].(string),
		usConf["dbname"].(string))
	db, err := sql.Open("postgres", connStr)
	return db, err
}


