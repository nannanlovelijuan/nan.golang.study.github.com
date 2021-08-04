package sql

import "testing"

func TestSelect(t *testing.T) {

	if err := Ping(); err != nil {
		t.Errorf("Ping() error,%s", err)
	}
}
