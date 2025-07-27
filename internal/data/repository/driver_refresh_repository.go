package repository

import "auth_service_medods/internal/domain/models"

type DriverRefreshRepository struct {
}

func (repo DriverRefreshRepository) Create(*models.Refresh) (*models.Refresh, error)
func (repo DriverRefreshRepository) Get(token string) (*models.Refresh, error)
func (repo DriverRefreshRepository) Delete(*models.Refresh) (*models.Refresh, error)
func (repo DriverRefreshRepository) Update(*models.Refresh) (*models.Refresh, error)
func (repo DriverRefreshRepository) CheckByUserUUID(user_uuid string) (bool, error)
func (repo DriverRefreshRepository) DeleteByUser(user_uuid string) (bool, error)
