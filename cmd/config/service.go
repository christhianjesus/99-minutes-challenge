package config

type Service struct {
	SrvHost string `env:"HOST,default=0.0.0.0"`
	SrvPort string `env:"PORT,default=8080"`
}

func (s Service) Addr() string {
	return s.SrvHost + ":" + s.SrvPort
}
