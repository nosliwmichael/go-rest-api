package configuration

import (
	"log"
	"time"

	"github.com/nosliwmichael/go-rest-api/pkg/configuration"
)

type (
	Config struct {
		AppName    string `yaml:"app-name"`
		HttpServer struct {
			Address      string        `yaml:"address"`
			ReadTimeout  time.Duration `yaml:"read-timeout"`
			WriteTimeout time.Duration `yaml:"write-timeout"`
		} `yaml:"http-server"`
		Api struct {
			ContextRoot string `yaml:"context-root"`
			Endpoints   struct {
				User     string `yaml:"user"`
				UserById string `yaml:"user-by-id"`
			} `yaml:"endpoints"`
		} `yaml:"api"`
		Database struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
			SID  string `yaml:"sid"`
		} `yaml:"database"`
	}
)

func LoadConfigs() (c *Config) {
	err := configuration.ReadConfigurations(&c, "./app_configs/app.yml")
	if err != nil {
		log.Println(err)
	}
	return
}
