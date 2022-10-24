package conf

type Config struct {
	Addr string

	DB struct {
		URI string
	}
}

func LoadFromEnv() (err error, conf Config) {

	return nil, Config{}
}
