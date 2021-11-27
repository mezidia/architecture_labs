package menu

import (
	"database/sql"
	"math"

	db_funcs "github.com/mezidia/architecture_labs/tree/main/docs/Lab3/server/db"
)

type Dish struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	Id     int   `json:"id"`
	Table  int   `json:"table_id"`
	Dishes []int `json:"dishes"`
}

type Price struct {
	Price float64
}

type Store struct {
	Db *sql.DB
}

var vatPercent = 7.0
var tipPercent = 3.0

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

func (s *Store) CreateOrder(id int, table int, dishes []int) error {
	var sum float64
	for _, v := range dishes {
		row := db_funcs.SelectOnePriceDishByID(s.Db, v)
		pr := Price{}
		err := row.Scan(&pr.Price)
		if err != nil {
			panic(err)
		}
		sum += pr.Price
	}

	var sumNoVat float64
	var tip float64

	sumNoVat = GetSumNoVat(sum)
	tip = GetTip(sum)

	db_funcs.InsertOneOrder(s.Db, id, table, dishes, sum, sumNoVat, tip)
	return nil
}

func GetSumNoVat(sumNoVat float64) (sum float64) {
	pricePercent := (sumNoVat / 100) * vatPercent
	sum = sumNoVat - pricePercent

	return math.Round(sum*100) / 100
}

func GetTip(tip float64) (sum float64) {
	tipPercent := (tip / 100) * tipPercent

	return math.Round(tipPercent*100) / 100
}

// func RoundPrice(float64) (sum float64) {
// 	roundedSum := math.Round(sum*100) / 100
// 	return roundedSum
// }
