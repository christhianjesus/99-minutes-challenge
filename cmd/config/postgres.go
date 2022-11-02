package config

import "fmt"

type DBConfig struct {
	DBHost     string `env:"DB_HOST,default=db"`
	DBPort     string `env:"DB_PORT,default=5432"`
	DBDatabase string `env:"DB_NAME,default=postgres"`
	DBUser     string `env:"DB_USER,default=postgres"`
	DBPassword string `env:"DB_PASSWORD,default=postgres"`
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBDatabase,
	)
}
