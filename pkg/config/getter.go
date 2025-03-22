package config

import (
	"github.com/The-EpaG/Go-BasicServer/config"
	"github.com/spf13/viper"
)

func Get() config.Config {
	return configurations
}

func ReadConfig(path, filename, format string) (*config.Config, error) {
	cfg := new(config.Config)
	viper.SetConfigName(filename)
	viper.SetConfigType(format)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
