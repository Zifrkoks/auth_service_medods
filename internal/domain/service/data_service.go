package service

import (
	"auth_service_medods/internal/data/repository"
	"auth_service_medods/internal/domain/models"
)

type DataService struct {
	users repository.UserRepository
}

func NewDataService(users repository.UserRepository) DataService {
	return DataService{users: users}
}
func (service DataService) GetUserData(user_id string) (user *models.User, err error) {
	return service.users.GetByGUID(user_id)
}
