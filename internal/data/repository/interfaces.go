package repository

import "auth_service_medods/internal/domain/models"

type (
	UserRepository interface {
		Create(*models.User) error
		GetByGUID(string) (*models.User, error)
	}
	RefreshRepository interface {
		Create(*models.Refresh) error
		Get(token string) (*models.Refresh, error)
		UserRefreshCount(user_uuid string) (int, error)
		DeleteByUser(user_uuid string) error
	}
)
