package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"weather-cache/internal/constants"
)

var Module = fx.Options(
	fx.Provide(Load),
)

func Load() (*Config, error) {
	viper.AutomaticEnv()

	path := viper.GetString(constants.ConfigPathKey)
	if path == "" {
		path = constants.DefaultConfigPath
	}

	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
