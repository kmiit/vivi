package types

type ServerConfig struct {
	Address	string
	Port 	int
}

type Config struct {
	Server ServerConfig
}