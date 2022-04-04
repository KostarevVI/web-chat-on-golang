// Package chatserver as main server app
package chatserver

import "web-chat-on-golang/dbstorage"

// Config chatserver struct
type Config struct {
	BindAddr string `toml:"bind_addr"`
	Database *dbstorage.Config
}

// NewConfig for chatserver
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Database: dbstorage.NewConfig(),
	}
}
