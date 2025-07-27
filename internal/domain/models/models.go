package models

import (
	"time"
)

type (
	User struct {
		GUID string
	}

	Refresh struct {
		TokenHash   string
		UserAgent   string
		Ip          string
		User        *User
		TimeCreated time.Time
	}
	AuthTokens struct {
		Jwt     string
		Refresh string
	}
	RefreshData struct {
		Jwt       string
		Refresh   string
		UserAgent string
		Ip        string
	}
	AuthData struct {
		Id        string
		UserAgent string
		Ip        string
	}
)
