package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	Timestamp	int32
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
	fmt.Printf("%+v", event)
	if !(event.Message != "" && event.URL != "") {
		http.Error(w, "Invalid Params", 400)
		return
	}

	event.Timestamp = int32(time.Now().Unix())
	events = append(events, event)
	json.NewEncoder(w).Encode(event)
}

func CreateEventEndpointPreFlight(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",	"Content-Type")
	return
}

func main() {
	events = append(events, Event{ Timestamp: int32(time.Now().Unix()), Message: "test msg", URL: "test url", LineNumber: "test lineNo", ColumnNumber: "test columnNo", ErrorObject: "test error", })
	events = append(events, Event{ Timestamp: int32(time.Now().Unix()), Message: "test 1 msg", URL: "test 1 url", LineNumber: "test 1 lineNo", ColumnNumber: "test 1 columnNo", ErrorObject: "test 1 error", })

	router := mux.NewRouter()

	router.HandleFunc("/events", GetEventsEndpoint).Methods("GET")
	router.HandleFunc("/event", CreateEventEndpoint).Methods("POST")
	router.HandleFunc("/event", CreateEventEndpointPreFlight).Methods("OPTIONS")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":12345", router))
}
