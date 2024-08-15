package config

type Config struct {
	ServerAddr string
}

func NewConfig() Config {
	return Config{
		ServerAddr: ":8080",
	}
}
