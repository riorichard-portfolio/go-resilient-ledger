package envconfig

import (
	"go-resilient-ledger/internal/infra/config"
	"go-resilient-ledger/internal/infra/validator"

	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"
)

func NewConfig(
	envPath string,
	validator *validator.Validator,
) (*config.Config, error) {
	v := viper.New()
	v.SetConfigFile(envPath)
	v.SetConfigType("env")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed read in config in New env config")
	}

	var envcfg EnvConfig
	err = v.Unmarshal(&envcfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshall from viper env config in New env config")
	}

	err = validator.Validate.Struct(&envcfg)
	if err != nil {
		return nil, errors.Wrap(err, "validation error env config New env config")
	}

	return &config.Config{
		PostgreConfig: config.PostgreConfig{
			URL: envcfg.PostgreURL,
		},
		AppConfig: config.AppConfig{
			Port:       envcfg.Port,
			TimeoutDur: envcfg.TimeoutDur,
		},
	}, nil
}
