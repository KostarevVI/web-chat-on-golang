// Package dbstorage
package dbstorage

// Config database structure
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig for DbStorage
func NewConfig() *Config {
	return &Config{}
}
