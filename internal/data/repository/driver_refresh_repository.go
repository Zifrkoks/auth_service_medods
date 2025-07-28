package repository

import (
	"auth_service_medods/internal/domain/models"
	"database/sql"
)

type DriverRefreshRepository struct {
	db *sql.DB
}

func NewDriverRefreshRepository(db *sql.DB) DriverRefreshRepository {
	return DriverRefreshRepository{db: db}
}
func (repo DriverRefreshRepository) Create(entity *models.Refresh) (err error) {
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
	data := repo.db.QueryRow("select token_hash,ip,user_id,user_agent,time_created from refreshes where token_hash = ?", token)
	err = data.Scan(&entity.TokenHash, &entity.Ip, &entity.User.GUID, &entity.UserAgent)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return
}

func (repo DriverRefreshRepository) DeleteByUser(user_uuid string) (err error) {
	_, err = repo.db.Exec("delete from refreshes where user_id = ?", user_uuid)
	return err
}
func (repo DriverRefreshRepository) UserRefreshCount(user_uuid string) (count int, err error) {
	err = repo.db.QueryRow("select count(*) from refreshes where user_id = ?", user_uuid).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return
}
