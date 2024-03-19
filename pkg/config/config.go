package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HttpAddress   string        `mapstructure:"HTTP_ADDRESS"`
	DbAddress     string        `mapstructure:"DB_ADDRESS"`
	Env           string        `mapstructure:"ENV"`
	GrpcPort      int           `mapstructure:"GRPC_PORT"`
	ClientAddress string        `mapstructure:"CLIENT_ADDRESS"`
	ClientTimeout time.Duration `mapstructure:"CLIENT_TIMEOUT"`
	ClientRetries int           `mapstructure:"CLIENT_RETRIES"`
}

func LoadConfig(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}
	}

	return cfg
}
