package config

type SRVConfig struct {
	SRVHost string `env:"HOST,default=0.0.0.0"`
	SRVPort string `env:"PORT,default=8080"`
}

func (c SRVConfig) Addr() string {
	return c.SRVHost + ":" + c.SRVPort
}
