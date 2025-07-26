package database

import (
	conf "auth_service_medods/config"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Config.DB.Host, conf.Config.DB.Port, conf.Config.DB.User, conf.Config.DB.Pass, conf.Config.DB.Name)
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

}
