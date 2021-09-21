package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeStruct struct {
	Time string `json:"time"`
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Visit /time page")
	fmt.Println("Endpoint Hit: indexPage")
}

func timePage(w http.ResponseWriter, r *http.Request) {
	timeValue := time.Now()
	timeData := TimeStruct{Time: timeValue.Format(time.RFC3339)}
	json.NewEncoder(w).Encode(timeData)
	fmt.Println("Endpoint Hit: timePage")
}

func routeRequests() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/time", timePage)
	http.ListenAndServe(":8795", nil)
}

func main() {
	routeRequests()
}
