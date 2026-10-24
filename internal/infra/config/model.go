package config

import "time"

type PostgreConfig struct {
	URL string
}

type AppConfig struct {
	Port      int
	TimeoutDur time.Duration
}

type Config struct {
	PostgreConfig PostgreConfig
	AppConfig     AppConfig
}
