package config

type Config struct {
	AdminToken string `env:"ADMIN_TOKEN"`

	SRVConfig
	DBConfig
}
