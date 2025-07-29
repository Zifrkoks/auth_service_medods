package authservicemedods

import (
	"os"
	"strconv"
	"time"
)

type BaseConfig struct {
	Auth   AuthConfig
	DB     DBConfig
	Server ServerConfig
}

type AuthConfig struct {
	JwtSecret string
	Duration  time.Duration
	WebHookIp string
}

type DBConfig struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

type ServerConfig struct {
	Host string
	Port string
}

var Config BaseConfig = BaseConfig{DB: load_db_conf_env(), Auth: load_auth_conf_env(), Server: load_server_conf_env()}

func load_db_conf_env() (db DBConfig) {
	if os.Getenv("DATABASE_PORT") == "" {
		db.Port = 5432
	} else {
		port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
		if err != nil {
			panic("DATABASE_PORT must be integer")
		}
		db.Port = port
	}

	db.Name = os.Getenv("DATABASE_NAME")
	if db.Host == "" {
		db.Host = "test_db"
	}
	db.Host = os.Getenv("DATABASE_HOST")
	if db.Host == "" {
		db.Host = "localhost"
	}
	db.User = os.Getenv("DATABASE_USER")
	if db.User == "" {
		db.User = "test_user"
	}
	db.Pass = os.Getenv("DATABASE_PASS")
	if db.Pass == "" {
		db.Pass = "test_pass"
	}
	return
}
func load_auth_conf_env() (auth AuthConfig) {
	if os.Getenv("AUTH_DURATION_MINUTES") == "" {
		auth.Duration = time.Minute * time.Duration(5)
	} else {
		minutes, err := strconv.Atoi(os.Getenv("AUTH_DURATION_MINUTES"))
		if err != nil {
			panic("AUTH_DURATION_MINUTES must be integer")
		}
		auth.Duration = time.Minute * time.Duration(minutes)
	}

	auth.JwtSecret = os.Getenv("AUTH_SECRET")
	if auth.JwtSecret == "" {
		auth.JwtSecret = "test"
	}
	auth.WebHookIp = os.Getenv("AUTH_WEBHOOK_IP")
	if auth.WebHookIp == "" {
		auth.WebHookIp = "127.0.0.1"
	}
	return
}

func load_server_conf_env() (server ServerConfig) {
	server.Host = os.Getenv("SERVER_HOST")
	if server.Host == "" {
		server.Host = "localhost"
	}
	server.Port = os.Getenv("SERVER_PORT")
	if server.Port == "" {
		server.Port = "8080"
	} else {
		_, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
		if err != nil {
			panic("SERVER_PORT must be integer")
		}
	}

	return
}
