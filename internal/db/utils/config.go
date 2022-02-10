package utils

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBDriver string `mapstructure:"DB_Driver"`
	DBSource string `mapstructure:"DB_Source"`
	Port     string `mapstructure:"Port"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	return
}
