package service

import (
	conf "auth_service_medods/config"
	repository "auth_service_medods/internal/data/repository"
	models "auth_service_medods/internal/domain/models"
	"auth_service_medods/internal/domain/utils"
	"auth_service_medods/internal/logger"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users     repository.UserRepository
	refreshes repository.RefreshRepository
}

func NewAuthService(users repository.UserRepository, refreshes repository.RefreshRepository) AuthService {
	return AuthService{users: users, refreshes: refreshes}
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
	refresh_bytes, err := base64.StdEncoding.DecodeString(data.Refresh)
	if err != nil {
		return models.AuthTokens{}, err
	}
	data.Refresh = string(refresh_bytes)
	jwt_token, err := utils.ParseJWT(data.Jwt)
	if err != nil {
		return models.AuthTokens{}, err
	}
	createdAt, err := jwt_token.Claims.GetIssuedAt()
	if err != nil {
		return tokens, err
	}
	user_uuid, err := jwt_token.Claims.GetSubject()
	if err != nil {
		return models.AuthTokens{}, err
	}
	old_refreshes, err := service.refreshes.GetByUser(user_uuid)
	if err != nil {
		logger.Log("db error in AuthService RefreshAuthTokens:" + err.Error())
		return
	}
	var old_refresh *models.Refresh = nil
	for _, v := range old_refreshes {
		err := bcrypt.CompareHashAndPassword([]byte(v.TokenHash), []byte(data.Refresh))
		if err == nil {
			old_refresh = &v
		}
	}
	if old_refresh == nil {
		return models.AuthTokens{}, errors.New("not found")
	}
	if old_refresh.UserAgent != data.UserAgent {
		err = service.refreshes.DeleteByUser(old_refresh.User.GUID)
		if err != nil {
			return models.AuthTokens{}, err
		}
		return tokens, errors.New("user-agent invalid")
	}
	if old_refresh.Ip != data.Ip {
		service.sendInfoAboutAuth(old_refresh.User.GUID, data.Ip)
	}
	logger.LogImportant(createdAt.Time.String())
	logger.LogImportant(old_refresh.TimeCreated.String())

	if old_refresh.TimeCreated.Round(time.Second).Compare(createdAt.Time) != 0 {
		err := service.refreshes.DeleteByUser(old_refresh.User.GUID)
		if err != nil {
			return tokens, errors.New("pair of refresh and jwt is invalid and database error")
		}
		return tokens, errors.New("pair of refresh and jwt is invalid")
	}
	tokens, err = service.createTokensPair(old_refresh.User, data.UserAgent, data.Ip)
	if err != nil {
		return models.AuthTokens{}, err
	}

	return

}

func (service AuthService) Logout(user_uuid string) (isLogout bool, err error) {
	err = service.refreshes.DeleteByUser(user_uuid)
	if err != nil {
		isLogout = true
	} else {
		isLogout = false
	}
	return
}

func (service AuthService) ValidateJWT(jwt_string string) (claims *jwt.RegisteredClaims, err error) {
	user_jwt, err := utils.ParseJWT(jwt_string)
	if err != nil {
		return nil, err
	}
	if !user_jwt.Valid {
		return nil, errors.New("jwt invalid")
	}

	uuid_user, err := user_jwt.Claims.GetSubject()
	if err != nil {
		return nil, errors.New("cant get subject")
	}
	count, err := service.refreshes.UserRefreshCount(uuid_user)
	logger.Log(strconv.Itoa(count))
	if count == 0 {
		return nil, errors.New("unauthorized by ValidateJWT")
	}
	if err != nil {
		return nil, err
	}
	if claims, ok := user_jwt.Claims.(*jwt.RegisteredClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("claims not created")
	}
}
func (service AuthService) createClaims(user models.User, createdAt time.Time) (claims jwt.RegisteredClaims) {
	claims.Subject = user.GUID
	claims.ExpiresAt = jwt.NewNumericDate(createdAt.Add((conf.Config.Auth.Duration)))
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
	err = service.refreshes.CreateOrUpdate(&refresh)
	if err != nil {
		return models.AuthTokens{}, err
	}
	tokens.Refresh = base64.StdEncoding.EncodeToString([]byte(tokens.Refresh))
	return
}
