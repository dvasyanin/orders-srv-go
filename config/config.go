package config

import (
	"github.com/micro/go-micro/util/log"
	"github.com/subosito/gotenv"
	"os"
)

type Config struct {
	App      *App
	Postgres *Postgres
}

func New() *Config {
	_ = gotenv.Load(".env")
	return &Config{
		App: &App{
			Port:    os.Getenv("PORT"),
			Version: os.Getenv("VERSION"),
		},
		Postgres: &Postgres{
			Addr:     os.Getenv("PG_ADDR"),
			Database: os.Getenv("PG_DB"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
		},
	}
}

type Postgres struct {
	Addr     string
	Database string
	User     string
	Password string
}

type App struct {
	Port    string
	Version string
}

func (c *Config) GetKey(key string) string {
	err := gotenv.Load(".env")
	if err != nil {
		log.Errorf("can't load env%+v", err)
	}
	return os.Getenv(key)
}
