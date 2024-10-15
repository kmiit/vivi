package types

type ServerConfig struct {
	Address string
	Port    int
}

type StorageConfig struct {
	WatchPath []string
}

type Config struct {
	Server  ServerConfig
	Storage StorageConfig
}
