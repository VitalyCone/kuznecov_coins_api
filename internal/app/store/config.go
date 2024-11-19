package store

type Config struct {
	DatabaseURL string
}

func NewConfig(DbURL string) (*Config) {
	return &Config{
		DatabaseURL: DbURL,
	}
}
