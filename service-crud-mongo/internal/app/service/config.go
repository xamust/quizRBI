package service

import (
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/cash"
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/store"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
	Cash     *cash.Config
}

func NewConfig() *Config {
	return &Config{}
}
