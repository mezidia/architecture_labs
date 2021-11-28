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

func (s *Store) CreateOrder(table int, dishes []int) error {
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

	sumNoVat := GetSumNoVat(sum)
	tip := GetTip(sum)

	db_funcs.InsertOneOrder(s.Db, table, dishes, sum, sumNoVat, tip)
	return nil
}

func GetSumNoVat(sum float64) (roundedPrice float64) {
	priceSum := (sum / 100) * vatPercent
	sumNoVat := sum - priceSum

	roundedPrice = RoundPrice(sumNoVat)

	return roundedPrice
}

func GetTip(sum float64) (roundedTip float64) {
	tipSum := (sum / 100) * tipPercent

	roundedTip = RoundPrice(tipSum)

	return roundedTip
}

func RoundPrice(sum float64) (roundedSum float64) {
	roundedSum = math.Round(sum*100) / 100

	return roundedSum
}
