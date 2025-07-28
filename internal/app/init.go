package app

import (
	conf "auth_service_medods/config"
	"auth_service_medods/internal/data/repository"
	"auth_service_medods/internal/domain/models"
	"auth_service_medods/internal/domain/service"
	"auth_service_medods/internal/logger"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	conf.Config.DB.Host, conf.Config.DB.Port, conf.Config.DB.User, conf.Config.DB.Pass, conf.Config.DB.Name)

var DB *sql.DB

func createTables() (err error) {
	_, err = DB.Exec("create table if not exists users(id varchar(40) primary key)")
	if err != nil {
		logger.Log("db:\ntable users not created. error:" + err.Error())
		return err
	}
	create_refreshes := "create table if not exists refreshes(" +
		"token_hash varchar(72) primary key, " +
		"ip varchar(16) not null, " +
		"user_id varchar(40) not null, " +
		"user_agent text not null, " +
		"time_created timestamp not null, " +
		"FOREIGN KEY (user_id) REFERENCES users (id) " +
		")"
	_, err = DB.Exec(create_refreshes)
	if err != nil {
		logger.Log("db:\ntable refreshes not created. error:" + err.Error())
	}
	return nil
}

func ConnectToDB() (err error) {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	err = db.Ping()
	if err != nil {
		return
	}
	DB = db
	createTables()
	return nil
}

func GetDB() *sql.DB {
	if DB == nil {
		if err := ConnectToDB(); err != nil {
			logger.Log("Reconnect failed: " + err.Error())
			return nil
		}
	}
	if err := DB.Ping(); err != nil {
		logger.Log("Reconnecting...")
		if newDB, err := sql.Open("postgres", psqlInfo); err == nil {
			DB = newDB
		}
	}

	return DB
}

func NewDriverUserRepository() repository.UserRepository {
	return repository.NewDriverUserRepository(GetDB())

}

func NewDriverRefreshRepository() repository.RefreshRepository {
	return repository.NewDriverRefreshRepository(GetDB())
}

func NewAuthService() service.AuthService {
	return service.NewAuthService(NewDriverUserRepository(), NewDriverRefreshRepository())
}
func NewDataService() service.DataService {
	return service.NewDataService(NewDriverUserRepository())
}

func InitUsers() (inited bool) {
	inited = true
	repo := NewDriverUserRepository()
	user1 := models.User{GUID: uuid.NewString()}
	err := repo.Create(&user1)
	if err != nil {
		logger.LogImportant("error:" + err.Error())
		inited = false
	} else {
		logger.LogImportant(user1.GUID)
	}
	user2 := models.User{GUID: uuid.NewString()}
	err = repo.Create(&user2)
	if err != nil {
		logger.LogImportant("error:" + err.Error())
		inited = false
	} else {
		logger.LogImportant(user2.GUID)
	}
	user3 := models.User{GUID: uuid.NewString()}
	err = repo.Create(&user3)
	if err != nil {
		logger.LogImportant("error:" + err.Error())
		inited = false
	} else {
		logger.LogImportant(user3.GUID)
	}
	return
}
