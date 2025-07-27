package database

import (
	conf "auth_service_medods/config"
	"auth_service_medods/internal/logger"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Config.DB.Host, conf.Config.DB.Port, conf.Config.DB.User, conf.Config.DB.Pass, conf.Config.DB.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec("create table users(id varchar(40) primary key)")
	if err != nil {
		logger.Log("db:\ntable users not created. error:" + err.Error())
	}
	create_refreshes := "create table refreshes(" +
		"token_hash varchar(72) primary key, " +
		"ip varchar(16) not null, " +
		"user_id varchar(40) not null, " +
		"user_agent text not null, " +
		"time_created timestamp nou null, " +
		"FOREIGN KEY (user_id) REFERENCES users (id) " +
		")"
	_, err = DB.Exec(create_refreshes)
	if err != nil {
		logger.Log("db:\ntable refreshes not created. error:" + err.Error())
	}
}
