package config

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

var appConfigs AppConfiguration

//AppConfiguration struct defining vars in _envs/
type AppConfiguration struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	SSL         bool   `mapstructure:"ssl"`
	ReleaseMode bool   `mapstructure:"release_mode"`
}

func initAppConfig() (*viper.Viper, error) {
	vc := viper.New()
	var err error
	var configFileName string

	vc.AutomaticEnv()
	processEnv := vc.GetString("ENV")

	if !(len(processEnv) > 0) {
		fmt.Println("No ENV received, defaulting to dev")
		configFileName = "env_dev"
	} else {
		configFileName = "env_" + processEnv
	}

	vc.SetConfigName(configFileName)     //e.g. env_dev.yaml
	vc.SetConfigType("yaml")             //config file type
	vc.AddConfigPath(path.Join("_envs")) //env file directory

	err = vc.ReadInConfig()

	return vc, err
}

//GetAppConfiguration returns pointer to appConfig
func GetAppConfiguration() *AppConfiguration {
	return &appConfigs
}

//NewAppConfiguration called by main() upon server init
func NewAppConfiguration() (*AppConfiguration, error) {
	vc, err := initAppConfig()
	if err != nil {
		return nil, fmt.Errorf("cannot read from config file: %s", err)
	}

	if err := vc.Unmarshal(&appConfigs); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s", err)
	}

	return &appConfigs, nil
}
