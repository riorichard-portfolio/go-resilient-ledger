package envconfig

import "time"

type EnvConfig struct {
	Port       int           `mapstructure:"PORT" validate:"required,min=3000"`
	PostgreURL string        `mapstructure:"POSTGRE_URL" validate:"required,url"`
	TimeoutDur time.Duration `mapstructure:"TIMEOUT_DURATION" validate:"required"`
}
