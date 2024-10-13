package configuration

import (
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

func ReadConfigurations[Config any](dest *Config, filePath string) error {
	ext := path.Ext(filePath)
	pref := strings.TrimSuffix(filePath, ext)
	profilePath := fmt.Sprintf("%s-%s%s", pref, getConfigProfileEnv(), ext)
	return loadConfigs(dest, filePath, profilePath)
}

func loadConfigs[Config any](dest Config, filePaths ...string) (err error) {
	for _, filePath := range filePaths {
		var data []byte
		if data, err = os.ReadFile(filePath); err != nil {
			break
		}
		if err = yaml.Unmarshal(data, dest); err != nil {
			break
		}
	}
	return err
}

func getConfigProfileEnv() string {
	rte := os.Getenv("CONFIG_PROFILE_ENV")
	if rte != "" {
		return strings.ToLower(rte)
	}
	return "dev"
}
