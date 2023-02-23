package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSource    string `mapstructure:"DB_SOURCE"`
	DBDriver    string `mapstructure:"DB_DRIVER"`
	Environment string `mapstructure:"ENVIRONMENT"`
	HTTPPort    string `mapstructure:"HTTP_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
