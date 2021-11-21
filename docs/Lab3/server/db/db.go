package db

import (
	"database/sql"
	"net/url"

	_ "github.com/lib/pq"
)

type Connection struct {
	DbName         string
	User, Password string
	Host           string
	DisableSSL     bool
}

func (c *Connection) ConnectionURL() string {
	dbUrl := &url.URL{
		Scheme: "postgres",
		Host:   c.Host,
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.DbName,
	}
	if c.DisableSSL {
		dbUrl.RawQuery = url.Values{
			"sslmode": []string{"disable"},
		}.Encode()
	}
	return dbUrl.String()
}

func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("postgres", c.ConnectionURL())
}

func (c *Connection) InsertOneDish(db *sql.DB, name string, price int) error {
	_, err := db.Exec("insert into dishes.dishes (name, price) values ($1, $2)",
		name, price)
	return err
}

func (c *Connection) InsertOneOrder(db *sql.DB, dishes []int, table int, suma1 float32, suma2 float32, suma3 float32) error {
	_, err := db.Exec("insert into orders.orders (dishes, table, suma1, suma2, suma3) values ($1, $2, $3, $4, $5)",
		dishes, table, suma1, suma2, suma3)
	return err
}

func (c *Connection) SelectAllDishes(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * from dishes.dishes")
	return rows, err
}

func (c *Connection) SelectAllOrders(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * from orders.orders")
	return rows, err
}
