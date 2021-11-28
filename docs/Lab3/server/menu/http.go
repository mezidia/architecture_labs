package menu

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mezidia/architecture_labs/tree/main/docs/Lab3/server/tools"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListMenu(store, rw)
		} else if r.Method == "POST" {
			handleCreateOrder(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListMenu(store *Store, rw http.ResponseWriter) {
	res, err := store.ListMenu()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleCreateOrder(r *http.Request, rw http.ResponseWriter, store *Store) {
	var c Order
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateOrder(c.Table, c.Dishes)
	if err == nil {
		tools.WriteJsonOk(rw, &c)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
