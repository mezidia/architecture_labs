package menu

import (
	"log"
	"net/http"

	"github.com/mezidia/architecture_labs/tree/main/docs/Lab3/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListMenu(store, rw)
		} else if r.Method == "POST" {

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
