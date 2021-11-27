package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorObject struct {
	Message string `json:"message"`
}

func WriteJsonBadRequest(rw http.ResponseWriter, message string) {
	writeJson(rw, http.StatusBadRequest, &errorObject{Message: message})
}

func WriteJsonInternalError(rw http.ResponseWriter) {
	writeJson(rw, http.StatusBadRequest, &errorObject{Message: "internal error happened"})
}

func WriteJsonOk(rw http.ResponseWriter, res interface{}) {
	writeJson(rw, http.StatusOK, res)
}

func writeJson(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
