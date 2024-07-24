package config

type OracleConfig struct {
	Url string
	Port int
	ServiceName string
	User string
	Password string
}

func LocalOracleConfig() OracleConfig {
	return OracleConfig{
		Url: "localhost",
		Port: 1521,
		ServiceName: "xe",
		User: "system",
		Password: "oracle",
	}
}