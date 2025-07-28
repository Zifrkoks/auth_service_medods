package repository

import (
	"auth_service_medods/internal/domain/models"
	"database/sql"
)

type DriverUserRepository struct {
	db *sql.DB
}

func NewDriverUserRepository(db *sql.DB) DriverUserRepository {
	return DriverUserRepository{db: db}
}
func (repo DriverUserRepository) Create(entity *models.User) (err error) {
	_, err = repo.db.Exec("insert into users (id) values ($1)", entity.GUID)
	return err
}
func (repo DriverUserRepository) GetByGUID(uuid string) (user *models.User, err error) {
	user = &models.User{}
	data := repo.db.QueryRow("select id from users where id = ?", uuid)
	err = data.Scan(&user.GUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return
}
