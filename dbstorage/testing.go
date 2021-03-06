package dbstorage

import (
	"fmt"
	"strings"
	"testing"
)

func TestBD(t *testing.T, databaseURL string) (*DbStorage, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	d := New(config)

	if err := d.Open(); err != nil {
		t.Fatal(err)
	}

	return d, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := d.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		if err := d.Close(); err != nil {
			return
		}
	}
}
