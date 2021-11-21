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

func SelectAllDishes(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * from dishes.dishes")
	return rows, err
}

func SelectAllOrders(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("select * from orders.orders")
	return rows, err
}

func (c *Connection) SelectOneDishByID(db *sql.DB, id int) (*sql.Rows, error) {
	row, err := db.Query("select * from dishes.dishes where id = $1", id)
	return row, err
}

func (c *Connection) SelectOneOrderByID(db *sql.DB, id int) (*sql.Rows, error) {
	row, err := db.Query("select * from orders.orders where id = $1", id)
	return row, err
}

func (c *Connection) UpdateDish(db *sql.DB, fieldName string, value int, id int) error {
	_, err := db.Exec("update dishes.dishes set "+fieldName+" = $1 where id = $2", value, id)
	return err
}

func (c *Connection) UpdateOrder(db *sql.DB, fieldName string, value, id int) error {
	_, err := db.Exec("update orders.orders set "+fieldName+" = $1 where id = $2", value, id)
	return err
}

func (c *Connection) DeleteDish(db *sql.DB, id int) error {
	_, err := db.Exec("delete from dishes.dishes where id = $1", id)
	return err
}

func (c *Connection) DeleteOrder(db *sql.DB, id int) error {
	_, err := db.Exec("delete from orders.orders where id = $1", id)
	return err
}
