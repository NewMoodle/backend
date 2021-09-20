package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

const (
	EnvLocal      = ".env.local"
	EnvProduction = ".env.production"
)

type Config struct {
	GinMode string
	Http    HTTPConfig
	Postgre PostgreConfig
	Auth    AuthConfig
}

type HTTPConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type PostgreConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
	Pools    int
}

type AuthConfig struct {
	PasswordHashCost int
}

func Load(env string) (Config, error) {
	var (
		cfg Config
		err error
	)

	if err = godotenv.Load(env); err != nil {
		return Config{}, err
	}

	cfg.GinMode = os.Getenv("GIN_MODE")

	cfg.Http, err = loadHttpConfig()
	if err != nil {
		return Config{}, err
	}

	cfg.Postgre, err = loadPostgreConfig()
	if err != nil {
		return Config{}, err
	}

	cfg.Auth, err = loadAuthConfig()
	if err != nil {
		return Config{}, err
	}

	return cfg, err
}

func loadHttpConfig() (HTTPConfig, error) {
	var (
		cfg HTTPConfig
		err error
	)

	cfg.Port, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return HTTPConfig{}, err
	}

	cfg.ReadTimeout, err = time.ParseDuration(os.Getenv("HTTP_TIMEOUT_READ"))
	if err != nil {
		return HTTPConfig{}, err
	}

	cfg.WriteTimeout, err = time.ParseDuration(os.Getenv("HTTP_TIMEOUT_WRITE"))
	if err != nil {
		return HTTPConfig{}, err
	}

	return cfg, err
}

func loadPostgreConfig() (PostgreConfig, error) {
	var (
		cfg PostgreConfig
		err error
	)

	cfg.Host = os.Getenv("POSTGRE_HOST")

	cfg.Port, err = strconv.Atoi(os.Getenv("POSTGRE_PORT"))
	if err != nil {
		return PostgreConfig{}, err
	}
	cfg.Username = os.Getenv("POSTGRE_USERNAME")
	cfg.Password = os.Getenv("POSTGRE_PASSWORD")
	cfg.Database = os.Getenv("POSTGRE_DATABASE")
	cfg.SSLMode = os.Getenv("POSTGRE_SSLMODE")
	cfg.Pools, err = strconv.Atoi(os.Getenv("POSTGRE_POOLS"))
	if err != nil {
		return PostgreConfig{}, err
	}

	return cfg, err
}

func loadAuthConfig() (AuthConfig, error) {
	var (
		cfg AuthConfig
		err error
	)

	cfg.PasswordHashCost, err = strconv.Atoi(os.Getenv("PASSWORD_HASH_COST"))
	if err != nil {
		return AuthConfig{}, err
	}

	return cfg, err
}
