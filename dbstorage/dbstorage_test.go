package dbstorage_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL_TEST")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=6559 dbname=web_chat_database sslmode=disable"
	}

	os.Exit(m.Run())
}
