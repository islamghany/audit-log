package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_SOURCE_NAME    string        `mapstructure:"DATABASE_SOURCE_NAME"`
	DB_MAX_OPEN_CONNECTION  int           `mapstructure:"DB_MAX_OPEN_CONNECTION"`
	DB_MAX_IDLE_CONNECTION  int           `mapstructure:"DB_MAX_IDLE_CONNECTION"`
	DB_MAX_IDLE_TIME        time.Duration `mapstructure:"DB_MAX_IDLE_TIME"`
	PORT                    int           `mapstructure:"PORT"`
	TOKEN_SYMMETRIC_KEY     string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	ACCESS_TOKEN_DURATION   time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_DURATION  time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	ACTIVATE_TOKEN_DURATION time.Duration `mapstructure:"ACTIVATE_TOKEN_DURATION"`
	MIGRATION_URL           string        `mapstructure:"MIGRATION_URL"`
}

func LoadENVFromFile(path, fileName, fileType string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
