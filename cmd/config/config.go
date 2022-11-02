package config

type Config struct {
	Service

	AdminToken string `env:"ADMIN_TOKEN"`
}
