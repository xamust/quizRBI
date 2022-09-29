package store

type Config struct {
	ConnString     string `toml:"conn_string"`
	DBName         string `toml:"db_name"`
	CollectionName string `toml:"collect_name"`
}

func NewConfig() *Config {
	return &Config{}
}
