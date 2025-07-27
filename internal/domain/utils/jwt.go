package utils

import (
	conf "auth_service_medods/config"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJWT(jwt_string string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(jwt_string, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Config.Auth.JwtSecret), nil
	})
}
