package db

import (
	"testing"
)

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://postgres:password@localhost/db_name?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}

func TestDbConnection_Open(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
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

func TestDbInsertDish(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	err := conn.InsertOneDish(db, "Apple", 72000)
	if err != nil {
		t.Error("Exec failed...")
	}
	defer db.Close()
}

func TestDbInsertOrder(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	err := conn.InsertOneOrder(db, []int{2, 3, 5, 7, 11, 13}, 3, 32., 1., 33.)
	if err != nil {
		t.Error("Exec failed...")
	}
	defer db.Close()
}

func TestDbSelectAllDishes(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	rows, err := conn.SelectAllDishes(db)
	if err != nil {
		t.Error("Exec failed...")
	}
	defer rows.Close()
}

func TestDbSelectAllOrders(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	rows, err := conn.SelectAllOrders(db)
	if err != nil {
		t.Error("Exec failed...")
	}
	defer rows.Close()
}

func TestDbSelectOneDish(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	rows, err := conn.SelectOneDishByID(db, 1)
	if err != nil {
		t.Error("Exec failed...")
	}
	defer rows.Close()
}

func TestDbSelectOneOrder(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	rows, err := conn.SelectOneOrderByID(db, 1)
	if err != nil {
		t.Error("Exec failed...")
	}
	defer rows.Close()
}

func TestDbUpdateDish(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	err := conn.UpdateDish(db, "price", 655000, 1)
	if err != nil {
		t.Error("Exec failed...")
	}
}

func TestDbUpdateOrder(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	err := conn.UpdateOrder(db, "table", 67000, 1)
	if err != nil {
		t.Error("Exec failed...")
	}
}

func TestDbDeleteDish(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	err := conn.DeleteDish(db, 1)
	if err != nil {
		t.Error("Exec failed...")
	}
}

func TestDbDeleteOrder(t *testing.T) {
	conn := &Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}

	db, _ := conn.Open()
	err := conn.DeleteOrder(db, 1)
	if err != nil {
		t.Error("Exec failed...")
	}
}
