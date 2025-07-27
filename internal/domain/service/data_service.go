package service

import "auth_service_medods/internal/data/repository"

type DataService struct {
	users repository.UserRepository
}
