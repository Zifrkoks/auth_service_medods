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
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		panic("DATABASE_PORT must be integer")
	}
	db.Name = os.Getenv("DATABASE_NAME")
	db.Host = os.Getenv("DATABASE_HOST")
	db.Port = port
	db.User = os.Getenv("DATABASE_USER")
	db.Pass = os.Getenv("DATABASE_PASS")
	return
}
func load_auth_conf_env() (auth AuthConfig) {
	minutes, err := strconv.Atoi(os.Getenv("AUTH_DURATION_MINUTES"))
	if err != nil {
		panic("AUTH_DURATION_MINUTES must be integer")
	}
	auth.Duration = time.Minute * time.Duration(minutes)
	auth.JwtSecret = os.Getenv("AUTH_SECRET")
	auth.WebHookIp = os.Getenv("AUTH_WEBHOOK_IP")
	return
}

func load_server_conf_env() (server ServerConfig) {
	server.Host = os.Getenv("SERVER_HOST")
	_, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic("SERVER_PORT must be integer")
	}
	server.Port = os.Getenv("SERVER_PORT")
	return
}
