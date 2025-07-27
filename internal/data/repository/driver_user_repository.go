package repository

import "auth_service_medods/internal/domain/models"

type DriverUserRepository struct {
}

func (repo DriverUserRepository) Create(*models.User) error
func (repo DriverUserRepository) GetByGUID(string) (*models.User, error)
