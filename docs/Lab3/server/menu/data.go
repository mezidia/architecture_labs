package menu

import (
	"database/sql"
	"fmt"
	"strconv"

	db_funcs "github.com/mezidia/architecture_labs/tree/main/docs/Lab3/server/db"
)

type Dish struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListMenu() ([]*Dish, error) {
	rows, err := db_funcs.SelectAllDishes(s.Db)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Dish
	for rows.Next() {
		var c Dish
		if err := rows.Scan(&c.Id, &c.Name, &c.Price); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Dish, 0)
	}
	return res, nil
}

var vatPercent = 7.0
var tipPercent = 3.0

func (s *Store) CreateOrder(id int, table int, dishes []int) error {
	// price, priceNoVAT and tip will be calculated automatically
	if (id <= 0) || (table <= 0) {
		return fmt.Errorf("something wrong with arguments")
	}

	var sum float64

	for _, dish := range dishes {
		var c Dish
		row, err := db_funcs.SelectOneDishByID(s.Db, dish)
		if err != nil {
			return err
		}
		if err := row.Scan(&c.Id, &c.Name, &c.Price); err != nil {
			return err
		}
		textPrice := fmt.Sprintf("%f", &c.Price)
		price, err := strconv.ParseFloat(textPrice, 64)
		sum = sum + price
	}

	var sumNoVat float64
	var tip float64

	sumNoVat = GetPercent(float64(vatPercent), sum)
	tip = GetPercent(float64(tipPercent), sum)

	err := db_funcs.InsertOneOrder(s.Db, id, table, dishes, sum, sumNoVat, tip)
	if err != nil {
		return err
	}
	return nil
}

func GetPercent(percent float64, price float64) (priceWithoutPercent float64) {
	priceWithoutPercent = (percent / price) * 100
	return priceWithoutPercent
}
