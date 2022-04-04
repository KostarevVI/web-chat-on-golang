// Package dbstorage
package dbstorage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// DbStorage structure
type DbStorage struct {
	config       *Config
	db           *sql.DB
	userDataRepo *UserDataRepo
}

// New DbStorage init
func New(config *Config) *DbStorage {
	return &DbStorage{
		config: config,
	}
}

// Open database connection
func (d *DbStorage) Open() error {
	db, err := sql.Open("postgres", d.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	d.db = db

	return nil
}

// Close database connection
func (d *DbStorage) Close() error {
	if err := d.db.Close(); err != nil {
		return err
	}
	return nil
}

// UserData implementation
func (d *DbStorage) UserData() *UserDataRepo {
	if d.userDataRepo != nil {
		return d.userDataRepo
	}

	d.userDataRepo = &UserDataRepo{
		dbstorage: d,
	}

	return d.userDataRepo
}
