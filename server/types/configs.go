package types

type DatabaseConfig struct {
	DbAddress  string
	DbPort     string
	DbPassword string
	DbNumber   int
}

type ServerConfig struct {
	Address string
	Port    string
}

type StorageConfig struct {
	WatchPath []string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Storage  StorageConfig
}
