package service

import (
	"auth_service_medods/internal/data/repository"
	"auth_service_medods/internal/domain/models"
)

type AuthService struct {
	users     repository.UserRepository
	refreshes repository.RefreshRepository
}

func (service AuthService) Auth(id string, user_agent string) (tokens models.AuthTokens)

func (service AuthService) RefreshAuthTokens(models.RefreshData) (tokens models.AuthTokens)

func (service AuthService) GetUserUUID(jwt_token string) (user_uuid string)

func (service AuthService) Logout(jwt_token string) bool
