package apiserver

type Config struct {
	ApiAddr string `toml:"bind_addr"`
}

func NewConfig(apiAddr string) (*Config) {
	return &Config{
		ApiAddr: apiAddr,
	}
}
