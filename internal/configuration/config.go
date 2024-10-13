package configuration

import (
	"log"

	"github.com/nosliwmichael/go-rest-api/pkg/configuration"
)

type (
	Config struct {
		AppName   string `yaml:"app-name"`
		Address   string `yaml:"address"`
		Endpoints struct {
			User       string `yaml:"user"`
			UserByName string `yaml:"user-by-name"`
		} `yaml:"endpoints"`
	}
)

func LoadConfigs() (c *Config) {
	err := configuration.ReadConfigurations(&c, "./app_configs/app.yml")
	if err != nil {
		log.Println(err)
	}
	return
}
