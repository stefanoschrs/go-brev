package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	Message      string `json:"msg"`
	URL          string `json:"url"`
	LineNumber   string `json:"lineNo"`
	ColumnNumber string `json:"columnNo"`
	ErrorObject  string `json:"error"`
}

var events []Event

func GetEventsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func CreateEventEndpoint(w http.ResponseWriter, req *http.Request) {
	var event Event

	_ = json.NewDecoder(req.Body).Decode(&event)
	if !(event.Message != "" && event.URL != "" && event.LineNumber != "" && event.ColumnNumber != "" && event.ErrorObject != "") {
		http.Error(w, "Invalid Params", 400)
		return
	}

	events = append(events, event)
	json.NewEncoder(w).Encode(event)
}

func main() {
	events = append(events, Event{
		Message:      "test msg",
		URL:          "test url",
		LineNumber:   "test lineNo",
		ColumnNumber: "test columnNo",
		ErrorObject:  "test error",
	})

	router := mux.NewRouter()

	router.HandleFunc("/events", GetEventsEndpoint).Methods("GET")
	router.HandleFunc("/event", CreateEventEndpoint).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":12345", router))
}
