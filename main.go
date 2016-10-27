package main

import (
	"encoding/json"
	"os"
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
	LineNumber   int `json:"lineNo"`
	ColumnNumber int `json:"columnNo"`
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	events = append(events, Event{ Message: "Uncaught ReferenceError: hello is not defined", URL: "https://0.0.0.0:12345/", LineNumber: 79, ColumnNumber: 4, ErrorObject: "{}", Timestamp: int32(time.Now().Unix()), })
	events = append(events, Event{ Message: "Uncaught ReferenceError: world is not defined", URL: "https://0.0.0.0:12345/", LineNumber: 89, ColumnNumber: 8, ErrorObject: "{}", Timestamp: int32(time.Now().Unix()), })

	router := mux.NewRouter()

	router.HandleFunc("/events", GetEventsEndpoint).Methods("GET")
	router.HandleFunc("/event", CreateEventEndpoint).Methods("POST")
	router.HandleFunc("/event", CreateEventEndpointPreFlight).Methods("OPTIONS")

	log.Println("Server listening on port " + port)
	err := http.ListenAndServeTLS(":" + port, "cert.pem", "cert.key", router)
	log.Fatal(err)
}
