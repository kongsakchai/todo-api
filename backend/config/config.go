package config

import (
	"os"
	"sync"
)

type Config struct {
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

		config = Config{
			DatabaseURL: databaseURL,
		}
	})

	return config
}
