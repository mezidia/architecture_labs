package db

import (
	"database/sql"
	"net/url"

	pq "github.com/lib/pq"
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

func SelectAllDishes(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM dishes")
	return rows, err
}

func SelectAllOrders(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM orders")
	return rows, err
}

func InsertOneDish(db *sql.DB, name string, price int) error {
	_, err := db.Exec("INSERT INTO dishes (name, price) VALUES ($1, $2)",
		name, price)
	return err
}

func InsertOneOrder(db *sql.DB, id int, table int, dishes []int, sum float64, sumNoVat float64, tip float64) error {

	_, err := db.Exec("INSERT INTO orders (id, table_id, dishes, order_sum, sum_no_vat, tip) VALUES ($1, $2, $3, $4, $5, $6)", id, table, pq.Array(dishes), sum, sumNoVat, tip)
	return err
}

func SelectOneDishByID(db *sql.DB, id int) (*sql.Rows, error) {
	row, err := db.Query("SELECT * FROM dishes WHERE id = $1", id)
	return row, err
}

func SelectOneOrderByID(db *sql.DB, id int) (*sql.Rows, error) {
	row, err := db.Query("SELECT * FROM orders WHERE id = $1", id)
	return row, err
}

func SelectOnePriceDishByID(db *sql.DB, id int) *sql.Row {
	row := db.QueryRow("SELECT price FROM dishes WHERE id = $1", id)
	return row
}

func UpdateDish(db *sql.DB, fieldName string, value int, id int) error {
	_, err := db.Exec("UPDATE dishes SET "+fieldName+" = $1 WHERE id = $2", value, id)
	return err
}

func UpdateOrder(db *sql.DB, fieldName string, value, id int) error {
	_, err := db.Exec("UPDATE orders SET "+fieldName+" = $1 WHERE id = $2", value, id)
	return err
}

func DeleteDish(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM dishes WHERE id = $1", id)
	return err
}

func DeleteOrder(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM orders WHERE id = $1", id)
	return err
}
