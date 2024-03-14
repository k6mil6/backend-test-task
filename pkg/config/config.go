package config

import "github.com/spf13/viper"

type Config struct {
	HttpAddress string `mapstructure:"HTTP_ADDRESS"`
	DbAddress   string `mapstructure:"DB_ADDRESS"`
	Env         string `mapstructure:"ENV"`
}

func LoadConfig(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("local")
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
