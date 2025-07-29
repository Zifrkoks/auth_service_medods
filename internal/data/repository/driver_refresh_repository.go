package repository

import (
	"auth_service_medods/internal/domain/models"
	"auth_service_medods/internal/logger"
	"database/sql"
	"fmt"
)

type DriverRefreshRepository struct {
	db *sql.DB
}

func NewDriverRefreshRepository(db *sql.DB) DriverRefreshRepository {
	return DriverRefreshRepository{db: db}
}
func (repo DriverRefreshRepository) CreateOrUpdate(entity *models.Refresh) (err error) {
	res, err := repo.db.Exec("update refreshes set token_hash = $1,time_created = $2 where user_id = $3 and user_agent = $4 and ip = $5", entity.TokenHash, entity.TimeCreated, entity.User.GUID, entity.UserAgent, entity.Ip)

	if err != nil {
		return err
	}
	edited, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if edited != 0 {
		return nil
	}
	_, err = repo.db.Exec(
		"insert into refreshes "+
			"(token_hash,"+
			"user_id,"+
			"user_agent,"+
			"ip,"+
			"time_created) values($1,$2,$3,$4,$5)",
		entity.TokenHash,
		entity.User.GUID,
		entity.UserAgent,
		entity.Ip,
		entity.TimeCreated)
	return
}
func (repo DriverRefreshRepository) Get(token string) (entity *models.Refresh, err error) {
	entity = &models.Refresh{User: &models.User{}}
	data := repo.db.QueryRow("select token_hash,ip,user_id,user_agent,time_created from refreshes where token_hash = $1", token)
	err = data.Scan(&entity.TokenHash, &entity.Ip, &entity.User.GUID, &entity.UserAgent, &entity.TimeCreated)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return
}

func (repo DriverRefreshRepository) DeleteByUser(user_uuid string) (err error) {
	_, err = repo.db.Exec("delete from refreshes where user_id = $1", user_uuid)
	return err
}
func (repo DriverRefreshRepository) UserRefreshCount(user_uuid string) (count int, err error) {
	err = repo.db.QueryRow("select count(*) from refreshes where user_id = $1", user_uuid).Scan(&count)
	if err == sql.ErrNoRows {
		logger.LogImportant("user refreshes not found")
		return 0, nil
	}
	return
}
func (repo DriverRefreshRepository) GetByUser(user_uuid string) (entities []models.Refresh, err error) {
	entities = []models.Refresh{}
	entity := models.Refresh{User: &models.User{}}
	data, err := repo.db.Query("select token_hash,ip,user_id,user_agent,time_created from refreshes where user_id = $1", user_uuid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return
	}
	defer data.Close()
	for data.Next() {
		err = data.Scan(&entity.TokenHash, &entity.Ip, &entity.User.GUID, &entity.UserAgent, &entity.TimeCreated)
		if err != nil {
			fmt.Println(err)
			continue
		}
		entities = append(entities, entity)
	}
	return
}
