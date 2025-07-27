package service

import (
	conf "auth_service_medods/config"
	repository "auth_service_medods/internal/data/repository"
	models "auth_service_medods/internal/domain/models"
	"auth_service_medods/internal/domain/utils"
	"auth_service_medods/internal/logger"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users     repository.UserRepository
	refreshes repository.RefreshRepository
}

func (service AuthService) Auth(data models.AuthData) (tokens models.AuthTokens, err error) {
	user, err := service.users.GetByGUID(data.Id)

	if err != nil {
		return models.AuthTokens{}, err
	}
	tokens, err = service.createTokensPair(user, data.UserAgent, data.Ip)
	if err != nil {
		return models.AuthTokens{}, err
	}
	return
}

func (service AuthService) RefreshAuthTokens(data models.RefreshData) (tokens models.AuthTokens, err error) {
	old_refresh, err := service.refreshes.Get(data.Refresh)
	if err != nil {
		logger.Log("db error in AuthService RefreshAuthTokens:" + err.Error())
		return
	}
	if old_refresh == nil {
		return tokens, errors.New("not found")
	}
	if old_refresh.UserAgent != data.UserAgent {
		err = service.refreshes.DeleteByUser(old_refresh.User.GUID)
		if err != nil {
			return models.AuthTokens{}, err
		}
		return tokens, errors.New("User-agent invalid")
	}
	if old_refresh.Ip != data.Ip {
		service.sendInfoAboutAuth(old_refresh.User.GUID, data.Ip)
	}
	jwt_token, err := utils.ParseJWT(data.Jwt)
	if err != nil {
		return models.AuthTokens{}, err
	}
	createdAt, err := jwt_token.Claims.GetIssuedAt()
	if old_refresh.TimeCreated.Compare(createdAt.Time) != 0 {
		service.refreshes.DeleteByUser(old_refresh.User.GUID)
		return tokens, errors.New("pair of refresh and jwt is invalid")
	}
	tokens, err = service.createTokensPair(old_refresh.User, data.UserAgent, data.Ip)
	if err != nil {
		return models.AuthTokens{}, err
	}
	return
}

func (service AuthService) Logout(jwt_token string) (isLogout bool, err error) {
	jwt, err := utils.ParseJWT(jwt_token)
	if err != nil {
		return false, errors.New("jwt is invalid")
	}
	user_uuid, err := jwt.Claims.GetSubject()
	if err != nil {
		return false, errors.New("jwt is invalid")
	}
	return false, service.refreshes.DeleteByUser(user_uuid)
}

func (service AuthService) createClaims(user models.User, createdAt time.Time) (claims jwt.RegisteredClaims) {
	claims.Subject = user.GUID
	claims.ExpiresAt = jwt.NewNumericDate(createdAt.Add((conf.Config.Auth.Duration * time.Minute)))
	claims.IssuedAt = jwt.NewNumericDate(createdAt)
	return
}

func (service AuthService) sendInfoAboutAuth(user_uuid string, ip string) {
	mes, err := json.Marshal(map[string]string{"mes": "user " + user_uuid + " try login from ip " + ip})
	if err != nil {
		logger.Log("error in marshal json")
	}
	_, err = http.Post(conf.Config.Auth.WebHookIp, "application/json", bytes.NewBuffer(mes))
	if err != nil {
		logger.Log("error in post request about auth")
	}
}

func (service AuthService) createTokensPair(user *models.User, userAgent string, ip string) (tokens models.AuthTokens, err error) {
	authTime := time.Now()
	tokens.Refresh = uuid.NewString()
	token_hash, err := bcrypt.GenerateFromPassword([]byte(tokens.Refresh), 10)
	if err != nil {
		return models.AuthTokens{}, err
	}
	claims := service.createClaims(*user, authTime)
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokens.Jwt, err = jwt.SignedString([]byte(conf.Config.Auth.JwtSecret))
	if err != nil {
		return models.AuthTokens{}, err
	}
	refresh := models.Refresh{}
	refresh.User = user
	refresh.TokenHash = string(token_hash)
	refresh.UserAgent = userAgent
	refresh.Ip = ip
	refresh.TimeCreated = authTime
	err = service.refreshes.Create(&refresh)
	if err != nil {
		return models.AuthTokens{}, err
	}
	return
}
