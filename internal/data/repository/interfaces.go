package repository

import "auth_service_medods/internal/domain/models"

type (
	UserRepository interface {
		Create(*models.User) error
		GetByGUID(string) (*models.User, error)
		GetAll() (entities []models.User, err error)
	}
	RefreshRepository interface {
		CreateOrUpdate(*models.Refresh) error
		Get(token string) (*models.Refresh, error)
		UserRefreshCount(user_uuid string) (int, error)
		DeleteByUser(user_uuid string) error
		GetByUser(user_uuid string) (entities []models.Refresh, err error)
	}
)
