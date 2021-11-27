package menu

import (
	"database/sql"

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
	//var tip float64

	sumNoVat = GetSumNoVat(sum)
	// tip = GetPercent(float64(tipPercent), sum)

	db_funcs.InsertOneOrder(s.Db, id, table, dishes, sum, sumNoVat, 50.0)
	return nil
	// 	if (id <= 0) || (table <= 0) {
	// 		return fmt.Errorf("something wrong with arguments")
	// 	}

	// 	var sum float64

	// 	for _, dish := range dishes {
	// 		var c Dish
	// 		row, err := db_funcs.SelectOneDishByID(s.Db, dish)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		for row.Next() {
	// 			if err := row.Scan(&c.Id, &c.Name, &c.Price); err != nil {
	// 				return err
	// 			}
	// 		}
	// 		textPrice := fmt.Sprintf("%f", &c.Price)
	// 		price, err := strconv.ParseFloat(textPrice, 64)
	// 		if err != nil {
	// 			return nil
	// 		}
	// 		sum = sum + price
	// 	}

	// 	var sumNoVat float64
	// 	var tip float64

	// 	sumNoVat = GetPercent(float64(vatPercent), sum)
	// 	tip = GetPercent(float64(tipPercent), sum)

	// 	err := db_funcs.InsertOneOrder(s.Db, id, table, dishes, sum, sumNoVat, tip)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }

}

func GetSumNoVat(sumNoVat float64) (sum float64) {
	pricePercent := (sumNoVat / 100) * vatPercent
	sum = sumNoVat - pricePercent
	return sum
}

// func GetTip(tip float64) (sum float64) {
// 	tipPercent := (tip / 100) * vatPercent
// 	sum = sumNoVat - pricePercent
// 	return sum
// }
