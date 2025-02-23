package config

import (
	"os"
	"sync"
)

type Config struct {
	Port        string
	DatabaseURL string
}

var once sync.Once
var config Config

func Get() Config {
	once.Do(func() {
		databaseURL := os.Getenv("DATABASE_URL")
		if len(databaseURL) == 0 {
			panic("DATABASE_URL is not set")
		}

		port := os.Getenv("PORT")
		if len(port) == 0 {
			port = "8080"
		}

		config = Config{
			DatabaseURL: databaseURL,
			Port:        port,
		}
	})

	return config
}
