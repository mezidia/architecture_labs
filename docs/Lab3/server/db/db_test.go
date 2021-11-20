package db

import (
	"testing"
)

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "Menu",
		User:       "postgres",
		Password:   "hogger",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://postgres:hogger@localhost/Menu?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}

func TestDbConnection_Open(t *testing.T) {
	conn := &Connection{
		DbName:     "Menu",
		User:       "postgres",
		Password:   "hogger",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, err := conn.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Error("Connection failed...")
	}
}
