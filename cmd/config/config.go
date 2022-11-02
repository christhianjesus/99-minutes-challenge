package config

type Config struct {
	AdminToken string `env:"ADMIN_TOKEN,default=token"`
	JWTSecret  string `env:"JWT_SECRET,default=secret"`

	SRVConfig
	DBConfig
}
