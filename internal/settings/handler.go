package settings

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Settings struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`

	Database struct {
		Name     string `env:"DB_NAME"`
		User     string `env:"DB_USER"`
		Port     string `env:"DB_PORT"`
		Host     string `env:"DB_HOST"`
		Password string `env:"DB_PASSWORD"`
	}
}

var settings Settings

func init() {
	if _, err := env.UnmarshalFromEnviron(&settings); err != nil {
		log.Fatal("Error on initialization envs: ", err)
	}
}

func GetSettings() Settings {
	return settings
}
