package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "db1",
		User:       "user1",
		Password:   "pass1",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://user1:pass1@localhost/db1?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
