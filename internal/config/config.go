package config

import (
	"fmt"
	"os"
)

type (
	Config struct {
		Server      Server
		Environment string
		AuthSalt    string
		Postgres    Postgres
	}
	Server struct {
		PORT string
	}
	Postgres struct {
		Host     string
		User     string
		Password string
		DbName   string
		Port     string
		SslMode  string
	}
)

func Init(configPath string) (Config, error) {
	var cfg = &Config{}
	setFromEnv(cfg)
	return *cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.Server.PORT = os.Getenv("SERVER_PORT")

	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.DbName = os.Getenv("POSTGRES_DBNAME")
	cfg.Postgres.SslMode = os.Getenv("POSTGRES_SSlMODE")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")

	cfg.AuthSalt = os.Getenv("AUTH_SALT")
}

func (p *Postgres) ToString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", p.Host, p.User, p.Password, p.DbName, p.Port, p.SslMode)
}
