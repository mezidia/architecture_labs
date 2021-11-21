package channels

import (
	"database/sql"

	datab "github.com/mezidia/architecture_labs/tree/main/docs/Lab3/server/db"
)

type Dish struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s Store) ListMenu() ([]*Dish, error) {
	conn := &datab.Connection{
		DbName:     "db_name",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		DisableSSL: true,
	}
	rows, err := conn.SelectAllDishes(s.Db)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Dish
	for rows.Next() {
		var c Dish
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Dish, 0)
	}
	return res, nil
}

// func (s *Store) CreateChannel(name string) error {
//     if len(name) < 0 {
//         return fmt.Errorf("channel name is not provided")
//     }
//     , err := s.Db.Exec("INSERT INTO channels (name) VALUES ($1)", name)
//     return err
// }
