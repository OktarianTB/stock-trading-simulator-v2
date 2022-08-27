package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The values are read by Viper from a config file or environment variables
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	TiingoAPI            string        `mapstructure:"TIINGO_API"`
	TiingoToken          string        `mapstructure:"TIINGO_TOKEN"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
}

// LoadConfig loads the configuration from a config file
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
