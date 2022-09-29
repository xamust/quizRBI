package cash

type Config struct {
	UpdateInterval int `toml:"update_interval"`
}

func NewConfig() *Config {
	return &Config{}
}
