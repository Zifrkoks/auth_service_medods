package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		GUID string
	}

	Refresh struct {
		TokenHash string
		UserAgent string
		User      *User
	}
	AuthTokens struct {
		Jwt     string
		Refresh string
	}
	RefreshData struct {
		Jwt       string
		Refresh   string
		UserAgent string
	}
)

func NewRefresh(user User, userAgent string) (refresh Refresh, refresh_token string, err error) {
	refresh_token = uuid.NewString()
	token_hash, err := bcrypt.GenerateFromPassword([]byte(refresh_token), 10)
	if err != nil {
		return
	}
	refresh.User = &user
	refresh.TokenHash = string(token_hash)
	refresh.UserAgent = userAgent

	return
}
