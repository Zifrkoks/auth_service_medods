package repository

import "auth_service_medods/internal/domain/models"

type (
	UserRepository interface {
		Create(models.User) (models.User, error)
		GetByGUID(string) (models.User, error)
	}
	RefreshRepository interface {
		Create(models.Refresh) (models.Refresh, error)
		Get(token string) (models.Refresh, error)
		Delete(models.Refresh) (models.Refresh, error)
		Update(models.Refresh) (models.Refresh, error)
	}
)
