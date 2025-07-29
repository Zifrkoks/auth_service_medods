package repository

import (
	"auth_service_medods/internal/domain/models"
	"database/sql"
	"fmt"
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
	data := repo.db.QueryRow("select id from users where id = $1", uuid)
	err = data.Scan(&user.GUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return
}
func (repo DriverUserRepository) GetAll() (entities []models.User, err error) {
	entities = []models.User{}
	entity := models.User{}
	data, err := repo.db.Query("select id from users")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return
	}
	defer data.Close()
	for data.Next() {
		err = data.Scan(&entity.GUID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		entities = append(entities, entity)
	}
	return
}
