package config

type Config struct {
	SrvHost    string `env:"HOST,default=0.0.0.0"`
	SrvPort    string `env:"PORT,default=8080"`
	AdminToken string `env:"ADMIN_TOKEN"`
}

func (c Config) Addr() string {
	return c.SrvHost + ":" + c.SrvPort
}
